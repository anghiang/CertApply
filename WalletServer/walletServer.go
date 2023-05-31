package main

import (
	"WalletClient/conf"
	"WalletClient/transact"
	"WalletClient/types"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/sessions"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"unsafe"
)

var sessionData map[string]interface{}

type WalletServer struct {
	port    uint16
	gateway string //区块链的节点地址
}
type TempTransactOpts struct {
	From      common.Address // Ethereum account to send the transaction from
	Nonce     *big.Int       // Nonce to use for the transaction execution (nil = use pending state)
	Value     *big.Int       // Funds to transfer along along the transaction (nil = 0 = no funds)
	GasPrice  *big.Int       // Gas price to use for the transaction execution (nil = gas price oracle)
	GasLimit  *big.Int       // Gas limit to set for the transaction execution (0 = estimate)
	Context   context.Context
	TempTrans *types.Transaction
}

type Account struct {
	PrivateKy string `json:"private_ky"`
	Address   string `json:"address"`
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

var store = sessions.NewCookieStore([]byte("yzg"))

func (ws *WalletServer) SetSignature(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//设置允许的方法
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")

		rawTxByte, err := ioutil.ReadAll(req.Body)
		fmt.Println("rawTxByte: ", rawTxByte)
		if err != nil {
			fmt.Println("ioutil.ReadAll err ", err)
			w.WriteHeader(http.StatusBadGateway)
		}
		var tx = new(types.Transaction)
		tx.UnmarshalJSON(rawTxByte)
		//err = json.Unmarshal(rawTxByte, &tx)
		//fmt.Println("tx :", tx)
		//if err != nil {
		//	fmt.Println("json.Unmarshal(rawTxByte,tx) err ", err)
		//	w.WriteHeader(http.StatusBadGateway)
		//}

		if privateKey, ok := sessionData["privateKey"]; ok {
			fmt.Println("InterfaceToString(privateKey): ", InterfaceToString(privateKey))

			signedTx, err := transact.SignTx(InterfaceToString(privateKey), tx)

			//peivateKey, err := crypto.ToECDSA(privateKetByte)
			//fmt.Println("===================,private", peivateKey)
			//if err != nil {
			//	fmt.Println("ToECDSA err ", err)
			//	return
			//}
			//signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, peivateKey)
			//if err != nil {
			//	fmt.Println("SignTx err ", err)
			//	return
			//}

			if err != nil {
				fmt.Println("SignTx err", err)
				w.WriteHeader(http.StatusBadGateway)
			}
			signedTxByte, err := signedTx.MarshalJSON()
			fmt.Println("==================signedTx: ", signedTx)
			if err != nil {
				fmt.Println("signedTx.MarshalJson() err ", err)
				w.WriteHeader(http.StatusBadGateway)
			}

			sessionID := "signedTxByte"
			sessionData[sessionID] = signedTxByte
		} else {
			fmt.Fprintln(w, "No session data found.")
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) GetSignature(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		if byteData, ok := sessionData["signedTxByte"]; ok {
			fmt.Println("banksignedTxByte: ", InterfaceToBytes(byteData))
			requestBody := bytes.NewBuffer(InterfaceToBytes(byteData))
			fmt.Println("requestBody: ", requestBody)
			resp, err := http.Post("http://localhost:8089/addCert", "application/json", requestBody)
			if err != nil {
				log.Println("HTTP request error:", err)

			}
			defer resp.Body.Close()
			fmt.Fprintln(writer, "Session data:", byteData)
		} else {
			fmt.Fprintln(writer, "No session data found.")
		}

	default:
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) GetTransSignature(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		if byteData, ok := sessionData["signedTxByte"]; ok {
			fmt.Println("banksignedTxByte: ", InterfaceToBytes(byteData))
			requestBody := bytes.NewBuffer(InterfaceToBytes(byteData))
			fmt.Println("requestBody: ", requestBody)
			resp, err := http.Post("http://localhost:8089/addTranscript", "application/json", requestBody)
			if err != nil {
				log.Println("HTTP request error:", err)

			}
			defer resp.Body.Close()
			fmt.Fprintln(writer, "Session data:", byteData)
		} else {
			fmt.Fprintln(writer, "No session data found.")
		}

	default:
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) PrivateKeyToAddress(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		privateKey := request.FormValue("privateKey_")
		ecdsaPri, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			return
		}
		address := crypto.PubkeyToAddress(ecdsaPri.PublicKey).Hex()
		fmt.Println(address)
		account := Account{
			PrivateKy: privateKey,
			Address:   address,
		}
		m, _ := json.Marshal(account)
		session, _ := store.Get(request, "PrivateKeySession")
		session.Values["privateKey"] = privateKey
		session.Save(request, writer)
		io.WriteString(writer, string(m[:]))
	default:
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) PemToPrivateKeyAndAddress(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		// 解析请求中的文件
		pemFile, handler, err := request.FormFile("file")
		if err != nil {
			http.Error(writer, "Failed to read file", http.StatusBadRequest)
			return
		}
		defer pemFile.Close()
		// 创建一个新文件来保存上传的文件内容
		dst, err := os.Create(handler.Filename)
		if err != nil {
			http.Error(writer, "Failed to create file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// 将文件内容复制到新文件中
		_, err = io.Copy(dst, pemFile)
		if err != nil {
			http.Error(writer, "Failed to copy file", http.StatusInternalServerError)
			return
		}
		// 获取保存文件的路径字符串
		filePath, err := os.Getwd() // 获取当前工作目录
		if err != nil {
			http.Error(writer, "Failed to get current working directory", http.StatusInternalServerError)
			return
		}
		privateKeyFromPEM, _, err := conf.LoadECPrivateKeyFromPEM(filePath + "/" + handler.Filename)
		if err != nil {
			writer.WriteHeader(http.StatusBadGateway)
			log.Println(err)
			return
		}
		//privateKey, _ := crypto.ToECDSA(privateKeyFromPEM)
		privateKeyStr := common.Bytes2Hex(privateKeyFromPEM)
		ecdsaPri, err := crypto.HexToECDSA(privateKeyStr)
		if err != nil {
			return
		}
		address := crypto.PubkeyToAddress(ecdsaPri.PublicKey).Hex()
		account := Account{
			PrivateKy: privateKeyStr,
			Address:   address,
		}
		sessionID := "privateKey"
		sessionData[sessionID] = privateKeyStr

		//session.Save(request, writer)
		m, _ := json.Marshal(account)
		io.WriteString(writer, string(m[:]))

		// 在文件操作完成后删除文件
		err = os.Remove(handler.Filename)
		if err != nil {
			//http.Error(writer, "Failed to delete file", http.StatusInternalServerError)
			fmt.Println(err)
		}
	default:
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func InterfaceToString(i interface{}) string {
	// 使用类型断言获取接口的具体值
	if str, ok := i.(string); ok {
		return str
	}

	// 如果接口的具体值不是字符串类型，根据实际需求进行转换
	// 这里只是一个示例，您可以根据实际情况进行处理
	return fmt.Sprintf("%v", i)
}

func InterfaceToBytes(i interface{}) []byte {
	// 使用类型断言获取接口的具体值
	if data, ok := i.([]byte); ok {
		return data
	}

	// 如果接口的具体值不是字节切片类型，根据实际需求进行转换
	// 这里只是一个示例，您可以根据实际情况进行处理

	// 使用反射获取接口的值
	value := reflect.ValueOf(i)

	// 如果接口的值是指针类型，则获取指针指向的值
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// 获取值的字节切片表示
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&value))
	return *(*[]byte)(unsafe.Pointer(&header))
}

func (ws *WalletServer) Run() {
	sessionData = make(map[string]interface{})
	http.HandleFunc("/setSignature", ws.SetSignature)
	http.HandleFunc("/getSignature", ws.GetSignature)
	http.HandleFunc("/stringSubmit", ws.PrivateKeyToAddress)
	http.HandleFunc("/pemSubmit", ws.PemToPrivateKeyAndAddress)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
