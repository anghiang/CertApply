package main

//func main() {
//r.GET("/download", func(c *gin.Context) {
//	cert := entity.Certs{Transcript: []interfaces.IBankCertificate{{
//		Id:             "1",
//		OwnerAddr:      common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
//		Name:           "czc",
//		IssueTime:      2023,
//		ValidityPeriod: 2043,
//	},
//		{
//			Id:             "2",
//			OwnerAddr:      common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
//			Name:           "czc2",
//			IssueTime:      2023,
//			ValidityPeriod: 2043,
//		},
//	}, HiddenData: []string{
//		"0x123",
//		"0x456",
//	}}
//	certJson, err := cert.MarshalJSON()
//	if err != nil {
//		panic(err)
//	}
//	certBytes, fileName := utils.NewFile(certJson)
//	c.Header("Content-Type", "application/octet-stream")
//	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
//	c.Writer.Write(certBytes)
//
//})
//r.POST("upload", func(c *gin.Context) {
//	file, _ := c.FormFile("aa")
//	f, err := file.Open()
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer f.Close()
//	data, _ := io.ReadAll(f)
//	var crt entity.Certs
//	err = crt.UnMarshalJSON(data)
//})
//http.HandleFunc("/download", func(writer http.ResponseWriter, request *http.Request) {
//
//	//cert := entity.NewCert("1", common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"), "czc", 2023, 2043)
//	cert := entity.Certs{Transcript: []entity.Cert{{
//		Id:             "1",
//		OwnerAddr:      common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
//		Name:           "czc",
//		IssueTime:      2023,
//		ValidityPeriod: 2043,
//	},
//		{
//			Id:             "2",
//			OwnerAddr:      common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
//			Name:           "czc2",
//			IssueTime:      2023,
//			ValidityPeriod: 2043,
//		},
//	}, HiddenData: []string{
//		"0x123",
//		"0x456",
//	}}
//	certJson, err := cert.MarshalJSON()
//	if err != nil {
//		panic(err)
//	}
//	certBytes, fileName := utils.NewFile(certJson)
//	writer.Header().Set("Content-Type", "application/octet-stream")
//	writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
//	writer.Write(certBytes)
//})
//http.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
//	values, _, err := request.FormFile("aa")
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusBadRequest)
//		return
//	}
//	defer values.Close() // 创建 multipart.Reader 来解析表单数据
//	reader := multipart.NewReader(request.Body, request.Header.Get("Content-Type"))
//
//	// 在 multipart.Reader 中查找 file 字段的部分
//	part, err := reader.NextPart()
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	// 创建一个缓冲区，将文件数据读入其中
//	buf := make([]byte, 1024)
//	n, err := part.Read(buf)
//	if err != nil && err != io.EOF {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	var c *entity.Certs
//	err = c.UnMarshalJSON(buf[:n])
//	fmt.Println("cert : ", c)
//})
//http.ListenAndServe("localhost:8082", nil)
//}
