package repository

import (
	"github.com/Deansquirrel/goToolMSSql2000"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/object"
	"github.com/kataras/iris/core/errors"
	"time"
)

const (
	sqlLoginVerify = "" +
		"SELECT COUNT(*) AS NUM " +
		"FROM [goreportlogininfo] " +
		"WHERE [loginname]=? and [loginpwd]=?"
	sqlGetMdIdByLogin = "" +
		"SELECT [MDID] " +
		"FROM [goreportlogininfo] " +
		"WHERE [loginname]=?"
)

type repZb struct {
	dbConfig *goToolMSSql2000.MSSqlConfig
}

func NewRepZb() *repZb {
	c := NewCommon()
	return &repZb{
		dbConfig: c.ConvertDbConfigTo2000(c.GetYwDbConfig()),
	}
}

//登录验证
func (r *repZb) LoginVerify(name string, pwd string) (bool, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlLoginVerify, name, pwd)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var num int
	num = -1
	for rows.Next() {
		err := rows.Scan(&num)
		if err != nil {
			return false, err
		}
	}
	if rows.Err() != nil {
		return false, rows.Err()
	}
	if num > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

//根据用户名获取门店ID
func (r *repZb) GetMdIdByLogin(name string) (int, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetMdIdByLogin, name)
	if err != nil {
		return -1, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var mdId int
	mdId = -1
	gotMdId := false
	for rows.Next() {
		err := rows.Scan(&mdId)
		if err != nil {
			return -1, err
		}
		gotMdId = true
	}
	if rows.Err() != nil {
		return -1, rows.Err()
	}
	if gotMdId {
		return mdId, nil
	} else {
		return -1, errors.New("got mdid failed")
	}
}

//获取门店报账收入日报数据
func (r *repZb) GetMdBaoZhShouR(mdId int, begDate time.Time, endDate time.Time) (*object.MdBaoZhShouRData, error) {
	//TODO
	return nil, nil
}
