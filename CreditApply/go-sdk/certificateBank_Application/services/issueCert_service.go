package services

import (
	"encoding/hex"
	"errors"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/contracts/implementation"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/entity"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"math/big"
)

// IssueCert 颁发证书
func IssueCert(_cert entity.Certs, certTgtHash [32]byte) error {

	//计算证书字段的TargetHash
	_cert.Cert.Metadata.TargetHash = hex.EncodeToString(certTgtHash[:])

	//将证书的TargetHash存放至合约上
	_, res, err := implementation.Issue(certTgtHash)
	if err != nil {
		return err
	}
	if res {
		_cert.Cert.AddCert()

	}
	return nil
}

// IssueCert 颁发证书
func IssueTranscript(_transcript entity.TranscriptCert, certTgtHash [32]byte) error {

	//计算证书字段的TargetHash
	_transcript.Transcript.Metadata.TargetHash = hex.EncodeToString(certTgtHash[:])

	//将证书的TargetHash存放至合约上
	_, res, err := implementation.Issue(certTgtHash)
	if err != nil {
		return err
	}
	if res {
		_transcript.Transcript.AddTranscript()

	}
	return nil
}

// VerifyCert 验证证书
func VerifyCert(_certByte []byte) (entity.Certs, error) {
	var cert entity.Certs
	err := cert.UnMarshalJSON(_certByte)
	if err != nil {
		return entity.Certs{}, err
	}
	//计算证书字段的TargetHash,以用于验证
	certTgtHash := cert.TargetHash()
	//将TargetHash与链上记录的进行匹配验证
	if implementation.IsIssued(certTgtHash) {
		//	//将证书文件转换为结构体，以便于在前台显示
		_err := cert.UnMarshalJSON(_certByte)
		if _err != nil {
			return entity.Certs{}, _err
		}
		return cert, nil
	}
	return entity.Certs{}, errors.New("无效的证书!")
}

// CertToBlockNum 获取已颁发证书对应的区块号及信息
func CertToBlockNum(_targetHash [32]byte) (*big.Int, *types.Block, error) {
	blkNum, err := implementation.CertToBlkNum(_targetHash)
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
