package models

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"
)

type Cert struct {
	Metadata Metadata
	Users    User
	Agencies Agency
	CertType int
}

func (c *Cert) AddCert() error {
	defer config.Db.Close()
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
	rows, err := config.Db.Query("select c.id, c.certNum,c.issueDate,c.validityPeriod,c.certName,a.mechanismName from cert as c inner join `users` as u on c.uid = u.id inner join agency as a on c.agencyId = a.id where u.address=? and c.certType=?", "0x5236A3B854232A845ae4AA7F56Df87B2377A9f34", c.CertType)
	if err != nil {
		return nil, err
	}

	var certSlice []Cert
	//scan:
	for rows.Next() {
		rows.Scan(&c.Metadata.Id, &c.Metadata.Number, &c.Metadata.IssueDate, &c.Metadata.ValidityPeriod, &c.Metadata.CertName, &c.Agencies.AgencyName)
		certSlice = append(certSlice, *c)
		//goto scan
	}
	return certSlice, nil
}
