package models

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"
)

// Transcript 成绩单
type Transcript struct {
	Metadata Metadata `json:"metadata"`
	Course   []Course `json:"course"`
	Users    User     `json:"users"`
	Agencies Agency   `json:"agencies"`
	CertType int      `json:"cert_type"`
}

// Course 分数
type Course struct {
	CourseName string
	Score      float64
}

func (t *Transcript) AddTranscript() error {
	t.CertType = 2
	uid, _ := QueryUidByAddress(t.Users.Address)
	res, err := config.Db.Exec("insert into cert(uid,certNum,issueDate,validityPeriod,agencyId,signature,certName,certType) values (?,?,?,?,?,?,?,?)", uid, t.Metadata.Number, t.Metadata.IssueDate, t.Metadata.ValidityPeriod, t.Agencies.Id, t.Metadata.Signature, t.Metadata.CertName, t.CertType)

	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()

	for _, course := range t.Course {
		_, err := config.Db.Exec("insert into course(courseName,score,uid,cert_id) values (?,?,?,?)", course.CourseName, course.Score, uid, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Transcript) QueryAllTranscripts(address string) (trans []Transcript, err error) {
	t.CertType = 2
	rows, err := config.Db.Query("SELECT ct.id,ct.certNum,ct.issueDate,ct.validityPeriod,ct.certName,a.mechanismName FROM cert ct INNER JOIN  `users` u ON ct.uid=u.id INNER JOIN agency a ON ct.agencyId=a.id WHERE u.address=? and ct.certType=?", address, t.CertType)
	if err != nil {
		return nil, err
	}
	//scan:
	for rows.Next() {
		rows.Scan(&t.Metadata.Id, &t.Metadata.Number, &t.Metadata.IssueDate, &t.Metadata.ValidityPeriod, &t.Metadata.CertName, &t.Agencies.AgencyName)
		trans = append(trans, *t)
		//goto scan
	}

	for _, tr := range trans {
		rows2, err := config.Db.Query("SELECT c.courseName,c.score FROM course c WHERE c.cert_id=?", tr.Metadata.Id)
		if err != nil {
			return nil, err
		}
		for rows2.Next() {
			c := new(Course)
			rows.Scan(&c.CourseName, &c.Score)
			tr.Course = append(tr.Course, *c)
			//goto scan
		}

	}
	return trans, nil
}
func (t *Transcript) QueryTranscriptByUser() ([]Transcript, error) {
	t.CertType = 2
	rows, err := config.Db.Query("select c.id, c.certNum,c.issueDate,c.validityPeriod,c.certName,a.mechanismName from cert as c inner join `users` as u on c.uid = u.id inner join agency as a on c.agencyId = a.id where u.address=? and c.certType = ?", t.Users.Address, t.CertType)
	if err != nil {
		return nil, err
	}
	var Transcripts []Transcript
	//scan:
	for rows.Next() {
		t.CertType = 2
		rows.Scan(&t.Metadata.Id, &t.Metadata.Number, &t.Metadata.IssueDate, &t.Metadata.ValidityPeriod, &t.Metadata.CertName, &t.Agencies.AgencyName)
		Transcripts = append(Transcripts, *t)
		//goto scan
	}
	return Transcripts, nil
}

func (t *Transcript) QueryTranscriptsById() error {
	rows := config.Db.QueryRow("select c.certNum Number,c.issueDate IssueDate,c.validityPeriod ValidityPeriod,c.certName CertName,c.signature Signature,a.mechanismName AgencyName,a.address Address,u.name UserName,u.address Address FROM cert as c inner join `users` as u on c.uid = u.id inner join agency as a on c.agencyId = a.id where  c.id = ?", t.Metadata.Id)
	//var certSlice *Cert
	err := rows.Scan(&t.Metadata.Number, &t.Metadata.IssueDate, &t.Metadata.ValidityPeriod, &t.Metadata.CertName, &t.Metadata.Signature, &t.Agencies.AgencyName, &t.Agencies.Address, &t.Users.UserName, &t.Users.Address)
	if err != nil {
		return err
	}

	fmt.Println("t.Metadata.Id: ", t.Metadata.Id)
	rows2, err := config.Db.Query("SELECT c.courseName,c.score FROM course c WHERE c.cert_id=?", t.Metadata.Id)
	if err != nil {
		return err
	}

	for rows2.Next() {
		var c Course
		rows2.Scan(&c.CourseName, &c.Score)
		t.Course = append(t.Course, c)
		//goto scan
	}

	return nil
}
