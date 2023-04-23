package services

//func NewCert() ([]byte, string, error) {
//var cert entity.Certs
//for _, v := range _certs {
//	cert.Transcript = append(cert.Transcript, v)
//}
//	cert := entity.CalHash()
//	certJson, err := cert.MarshalJSON()
//	if err != nil {
//		return nil, "", err
//	}
//
//	certBytes, fileName := utils.NewFile(certJson)
//	return certBytes, fileName, err
//}

//func VerifyIsCertExist(uAddr common.Address, certs []interfaces.IBankCertificate) bool {
//	chainCerts := implementation.GetCert(uAddr)
//	for _, v := range chainCerts {
//		for _, cert := range certs {
//			if reflect.DeepEqual(v, cert) {
//				return true
//			}
//		}
//	}
//	return false
//}
