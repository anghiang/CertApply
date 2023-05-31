package models

// Metadata 证书元数据
type Metadata struct {
	Id             int    `json:"id"`
	Number         string `json:"number"`
	IssueDate      string `json:"issue_date"`
	ValidityPeriod int    `json:"validity_period"`
	Signature      string `json:"signature"`
	CertName       string `json:"cert_name"`
	TargetHash     string `json:"target_hash"`
}
