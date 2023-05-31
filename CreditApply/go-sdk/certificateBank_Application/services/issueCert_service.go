package services

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/contracts/implementation"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/entity"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/utils"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"math/big"
)

// IssueCert 颁发证书
// IssueCert 颁发证书
func IssueCert(_cert entity.Certs, certTgtHash [32]byte) error {

	_cert.Cert.Metadata.TargetHash = hex.EncodeToString(certTgtHash[:])
	//将证书的TargetHash存放至合约上
	_, res, err := implementation.Issue(certTgtHash)
	if err != nil {
		return err
	}
	if res {
		signature, err := SignCertTgtHsh(certTgtHash, _cert.Cert.Agencies.Address)
		if err != nil {
			return err
		}
		_cert.Cert.Metadata.Signature = signature
		_cert.Cert.AddCert()
		_, err = _cert.MarshalJSON()
		if err != nil {
			return err
		}

	}
	return nil
}

func RevokeCert(_targetHash [32]byte) (error, int) {
	_, err, code := implementation.Revoke(_targetHash)
	if code == 400 || code == 401 || code == 402 {
		return err, code
	}

	return nil, code
}

// SignCertTgtHsh 证书颁发机构为证书的targetHash签名
func SignCertTgtHsh(_msg [32]byte, _address string) (string, error) {
	priKeyfile, err := utils.FindKeyFile(_address)
	if err != nil {
		return "", err
	}
	if priKeyfile != "" {
		priKey, err := utils.GetPrivateKey(priKeyfile)
		if err != nil {
			return "", err
		}
		signature, err := utils.SignMsg(_msg, priKey)
		if err != nil {
			return "", err
		}
		return hex.EncodeToString(signature), nil
	}
	return "", errors.New("签名失败")
}

// VerySignature 对已颁发证书证书的签名进行验证
func VerySignature(_targetHash, _signature string, _address string) (bool, error) {
	priKeyfile, err := utils.FindKeyFile(_address)
	if err != nil {
		return false, err
	}
	if priKeyfile != "" {
		priKey, err := utils.GetPrivateKey(priKeyfile)
		if err != nil {
			return false, err
		}
		targetHashByte, err := hex.DecodeString(_targetHash)
		var targetHashArray [32]byte
		copy(targetHashArray[:], targetHashByte)
		signatureByte, err := hex.DecodeString(_signature)
		verify := utils.VerifyMsg(targetHashArray, signatureByte, &priKey.PublicKey)
		if verify {
			return true, nil
		}
		return false, nil
	}
	return false, nil
}

// VerifyCert 验证证书
func VerifyCert(cert *entity.Certs) (bool, error) {
	//计算证书字段的TargetHash,以用于验证
	certTgtHash := cert.TargetHash()
	fmt.Println("certTgtHash: ", certTgtHash)
	//将TargetHash与链上记录的进行匹配验证
	if implementation.IsIssued(certTgtHash) && !implementation.IsRevoked(certTgtHash) {
		//	//将证书文件转换为结构体，以便于在前台显示
		return true, nil
	}
	return false, errors.New("无效的证书!")
}

// VerifyTrans 验证成绩单
func VerifyTrans(trans *entity.TranscriptCert) (bool, error) {
	//计算证书字段的TargetHash,以用于验证
	certTgtHash := trans.TranscriptTargetHash()
	//将TargetHash与链上记录的进行匹配验证
	if implementation.IsIssued(certTgtHash) && !implementation.IsRevoked(certTgtHash) {
		//	//将证书文件转换为结构体，以便于在前台显示
		return true, nil
	}
	return false, errors.New("无效的证书!")
}

// IssueTranscript 颁成绩证书
func IssueTranscript(_transcript entity.TranscriptCert, certTgtHash [32]byte) error {
	fmt.Println("_transcript: ", _transcript)
	//计算证书字段的TargetHash
	_transcript.Transcript.Metadata.TargetHash = hex.EncodeToString(certTgtHash[:])
	fmt.Println("_transcript.Transcript.Metadata.TargetHash: ", _transcript.Transcript.Metadata.TargetHash)

	//将证书的TargetHash存放至合约上
	_, res, err := implementation.Issue(certTgtHash)
	if err != nil {
		return err
	}
	if res {
		signature, err := SignCertTgtHsh(certTgtHash, _transcript.Transcript.Agencies.Address)
		if err != nil {
			return err
		}
		_transcript.Transcript.Metadata.Signature = signature
		_transcript.Transcript.AddTranscript()
	}
	return nil
}

// RevokeToBlockNum 获取已撤销证书对应的区块号及信息
func RevokeToBlockNum(_targetHash [32]byte) (*big.Int, *types.Block, error) {
	blkNum, err := implementation.RevokeToBlkNum(_targetHash)
	if err != nil {
		return nil, nil, err
	}
	//判断哈希对应的链上的块是否为0
	if blkNum.Cmp(big.NewInt(0)) == 0 {
		return nil, nil, errors.New("无效的证书!")
	}
	//通过证书对应的区块号获取区块信息
	block, b_err := implementation.GetBlockByNumber(blkNum.Int64())
	if b_err != nil {
		return nil, nil, b_err
	}

	return blkNum, block, nil

}
