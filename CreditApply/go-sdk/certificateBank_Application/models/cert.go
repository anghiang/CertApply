package models

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"
)

type Cert struct {
	Metadata Metadata `json:"metadata"`
	Users    User     `json:"users"`
	Agencies Agency   `json:"agencies"`
	CertType int      `json:"cert_type"`
}

func (c *Cert) AddCert() error {
	c.CertType = 1
	uid, err := QueryUidByAddress(c.Users.Address)
	if err != nil {
		panic(err)
	}

	res, err := config.Db.Exec("insert into cert(uid,certNum,issueDate,validityPeriod,agencyId,signature,certName,certType) values (?,?,?,?,?,?,?,?)", uid, c.Metadata.Number, c.Metadata.IssueDate, c.Metadata.ValidityPeriod, c.Agencies.Id, c.Metadata.Signature, c.Metadata.CertName, c.CertType)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	fmt.Println("res := ", id)
	return nil
}

func (c *Cert) QueryCertsByUser() ([]Cert, error) {
	c.CertType = 1
	fmt.Println(c.Users.Address)
	rows, err := config.Db.Query("select c.id, c.certNum,c.issueDate,c.validityPeriod,c.certName,a.mechanismName from cert as c inner join `users` as u on c.uid = u.id inner join agency as a on c.agencyId = a.id where u.address=? and c.certType = ?", c.Users.Address, c.CertType)
	if err != nil {
		return nil, err
	}

	var certSlice []Cert
	//scan:
	for rows.Next() {
		c.CertType = 1
		rows.Scan(&c.Metadata.Id, &c.Metadata.Number, &c.Metadata.IssueDate, &c.Metadata.ValidityPeriod, &c.Metadata.CertName, &c.Agencies.AgencyName)
		certSlice = append(certSlice, *c)
		//goto scan
	}
	return certSlice, nil
}

func (c *Cert) QueryCertsById() error {
	rows := config.Db.QueryRow("select c.certNum Number,c.issueDate IssueDate,c.validityPeriod ValidityPeriod,c.certName CertName,c.signature Signature,a.mechanismName AgencyName,a.address Address,u.name UserName,u.address Address FROM cert as c inner join `users` as u on c.uid = u.id inner join agency as a on c.agencyId = a.id where  c.id = ?", c.Metadata.Id)
	//var certSlice *Cert
	err := rows.Scan(&c.Metadata.Number, &c.Metadata.IssueDate, &c.Metadata.ValidityPeriod, &c.Metadata.CertName, &c.Metadata.Signature, &c.Agencies.AgencyName, &c.Agencies.Address, &c.Users.UserName, &c.Users.Address)
	if err != nil {
		return err
	}
	return nil
}
