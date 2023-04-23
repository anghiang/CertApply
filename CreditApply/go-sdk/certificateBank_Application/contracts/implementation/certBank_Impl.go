package implementation

//
//import (
//	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/abi"
//	"github.com/FISCO-BCOS/go-sdk/client"
//	"github.com/ethereum/go-ethereum/common"
//	"log"
//)
//
//var contractAddress = common.HexToAddress("0xcb465127073489235bbb8cf6995172a7c7420b3f")
//
//// 合约实例 instance
//var instance *abi.CertBank
//
//// 合约交易体
//var session *abi.CertBankSession
//
//func InitContract(client *client.Client) {
//	tempIns, err := abi.NewCertBank(contractAddress, client)
//	if err != nil {
//		log.Fatal(err)
//	}
//	tempSession := &abi.CertBankSession{Contract: tempIns, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
//	instance = tempIns
//	session = tempSession
//}
//func GetCert(uAddr common.Address) interface{} {
//	value, err := session.GetUserCertificate(uAddr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return value
//}
//func GetCredit(uAddr common.Address) interface{} {
//	value, err := session.GetUserCredits(uAddr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return value
//}
//func GetUserInfo(uAddr common.Address) interface{} {
//	value, err := session.GetUserByAddress(uAddr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return value
//
//}
