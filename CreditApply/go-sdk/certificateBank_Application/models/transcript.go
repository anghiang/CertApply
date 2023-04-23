package models

import (
	"github.com/FISCO-BCOS/go-sdk/certificateBank_Application/config"
)

// Transcript 成绩单
type Transcript struct {
	Metadata Metadata
	Course   []Course
	Users    User
	Agencies Agency
	CertType int
}

// Course 分数
type Course struct {
	CourseName string
	Score      float64
}

func (t *Transcript) AddTranscript(address string) error {
	t.CertType = 2
	uid, _ := QueryUidByAddress(address)

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
func (t *Transcript) QueryTranscriptsByUser(address string) (trans []Transcript, err error) {
	t.CertType = 2
	rows, err := config.Db.Query("SELECT ct.id,ct.certNum,ct.issueDate,ct.validityPeriod,ct.certName,a.mechanismName FROM cert ct INNER JOIN  `users` u ON ct.uid=u.id INNER JOIN agency a ON ct.agencyId=a.id WHERE u.address=? and ct.certType=?", address, t.CertType)
	if err != nil {
		return nil, err
	}
	defer config.Db.Close()
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

//func (t *Transcript)()  {
//
//}