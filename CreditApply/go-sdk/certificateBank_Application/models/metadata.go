package models

// Metadata 证书元数据
type Metadata struct {
	Id             int
	Number         string
	IssueDate      string
	ValidityPeriod int
	Signature      string
	CertName       string
	TargetHash     string
}
