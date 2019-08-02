package repository

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolMSSql2000"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/object"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

const (
	sqlLoginVerify = "" +
		"SELECT COUNT(*) AS NUM " +
		"FROM [goreportlogininfo] " +
		"WHERE [loginname]=? and [loginpwd]=?"
	sqlGetMdIdByLogin = "" +
		"SELECT [MDID] " +
		"FROM [goreportlogininfo] " +
		"WHERE [loginname]=?"
	sqlGetZzInfo = "" +
		"SELECT [transid],[transname] " +
		"FROM [jltytranst] " +
		"ORDER BY [transdesc],[transid]"
	sqlGetKzInfo = "" +
		"SELECT [gsid],[gsname] " +
		"FROM [jltygst] " +
		"WHERE RIGHT([gssign],1) = 3 " +
		"ORDER BY [gsid]"
	sqlGetQzInfo = "" +
		"SELECT [gsid],[gsname] " +
		"FROM [jltygst] " +
		"WHERE RIGHT([gssign],1) = 2 " +
		"ORDER BY [gsid]"
	sqlGetBaoZhShouRSummaryData = "" +
		"SELECT [xshsr],[xsxjsr] = floor(xsxjsr * xsxjdybl / 10) / 10 ,[xjrate] = 1,[xsszsr] = floor(xsszsr * xsxjdybl / 10) / 10,[szrate] = 1," +
		"[xsjycs] = ceiling(xsjycs * xsjycsdybl / 100),[csrate] = 1" +
		"FROM [ywmdxssrhzt_md] A " +
		"  INNER JOIN [xtmdyystatus_md] B ON A.[xsmdid] = B.[mdyyid] " +
		"    AND CONVERT(varchar(100), A.[xshsr], 112)= B.[mdyydate] " +
		"    AND B.[mdyysjtype] & 1 <> 0" +
		"WHERE [xshsr]>=? and [xshsr]<=? and [xsmdid] = ?"
	sqlGetBaoZhShouRZzDetailData = "" +
		"SELECT [xshsr],round([xszzje] * [xszzdybl] / 100,1) AS [zzje],[xszzid] AS [zzid] " +
		"FROM [ywmdxssrzzdt_md] A " +
		"INNER JOIN [xtmdyystatus_md] B ON A.[xsmdid] = B.[mdyyid] " +
		"  AND CONVERT(varchar(100), A.[xshsr], 112)= B.[mdyydate] " +
		"  AND B.[mdyysjtype] & 1 <> 0 " +
		"WHERE [xshsr]>=? and [xshsr]<=? and [xsmdid] = ?"
	sqlGetBaoZhShouRKzDetailData = "" +
		"SELECT [xshsr],round([xskzje] * [xskzdybl] / 100,1) AS [kzje],[xskzid] AS [kzid] " +
		"FROM [ywmdxssrkzmxhzt_md] A " +
		"INNER JOIN [xtmdyystatus_md] B ON A.[xsmdid] = B.[mdyyid] " +
		"  AND CONVERT(varchar(100), A.[xshsr], 112)= B.[mdyydate] " +
		"  AND B.[mdyysjtype] & 1 <> 0 " +
		"WHERE [xshsr] >= ? and [xshsr] <= ? and [xsmdid] = ?"
	sqlGetBaoZhShouRQzDetailData = "" +
		"SELECT [xshsr],round([xsqzje] * [xsqzdybl] / 100,1) AS [qzje],[xsqzid] AS [qzid] " +
		"FROM [ywmdxssrqzmxhzt_md] A " +
		"INNER JOIN [xtmdyystatus_md] B ON A.[xsmdid] = B.[mdyyid] " +
		"  AND CONVERT(varchar(100), A.[xshsr], 112)= B.[mdyydate] " +
		"  AND B.[mdyysjtype] & 1 <> 0 " +
		"WHERE [xshsr] >= ? and [xshsr] <= ? and [xsmdid] = ?"
	sqlGetMdName = "" +
		"SELECT [BRNAME] " +
		"FROM [JLTYBRANCHT] " +
		"WHERE [BRID] = ?"

	sqlGetDisableList = "" +
		"declare @mdid int " +
		"declare @pset nvarchar(255)  " +
		"select @mdid = ?  " +
		"select @pset = isnull(mdus.uvalue,us.uvalue) " +
		"from zlusersettings us " +
		"left join ( " +
		"    select ukey,uvalue " +
		"    from zlmdusersettings mdus " +
		"    inner join jltybrancht br on br.brripecode = mdus.brripecode " +
		"    where br.brid = @mdid " +
		") mdus on mdus.ukey = us.ukey where us.ukey = 'XsSrDyxSz'  " +
		"select rtrim(ltrim(dbo.fn_strpart(dbo.fn_strpart(@pset,'|',sn),'=',1))) as [n] " +
		"from xtstaticsn where sn > 0 and sn <= len(@pset)-len(replace(@pset,'|',''))+1"
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
			errMsg := fmt.Sprintf("LoginVerify read data err: %s", err.Error())
			log.Error(errMsg)
			return false, errors.New(errMsg)
		}
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("LoginVerify read data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return false, errors.New(errMsg)
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
			errMsg := fmt.Sprintf("GetMdIdByLogin read data err: %s", err.Error())
			log.Error(errMsg)
			return -1, errors.New(errMsg)
		}
		gotMdId = true
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("GetMdIdByLogin read data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return -1, errors.New(errMsg)
	}
	if gotMdId {
		return mdId, nil
	} else {
		return -1, errors.New("get mdid failed")
	}
}

//获取门店名称
func (r *repZb) GetMdName(mdId int) (string, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetMdName, mdId)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = rows.Close()
	}()
	var mdName string
	gotMdName := false
	for rows.Next() {
		err = rows.Scan(&mdName)
		if err != nil {
			errMsg := fmt.Sprintf("GetMdName read data err: %s", err.Error())
			log.Error(errMsg)
			return "", errors.New(errMsg)
		}
		gotMdName = true
		break
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("GetMdName read data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return "", errors.New(errMsg)
	}
	if gotMdName {
		return mdName, nil
	} else {
		return "", errors.New("get mdName failed")
	}
}

//获取转账方式 map[int]string ID-NAME
func (r *repZb) GetZzInfo() (map[int]string, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZzInfo)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var transId int
	var transName string
	rData := make(map[int]string)
	for rows.Next() {
		err = rows.Scan(&transId, &transName)
		if err != nil {
			errMsg := fmt.Sprintf("read zzinfo data err: %s", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rData[transId] = transName
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("read zzinfo data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rData, nil
}

//获取卡种 map[int]string ID-NAME
func (r *repZb) GetKzInfo() (map[int]string, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetKzInfo)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var gsId int
	var gsName string
	rData := make(map[int]string)
	for rows.Next() {
		err = rows.Scan(&gsId, &gsName)
		if err != nil {
			errMsg := fmt.Sprintf("read kzinfo data err: %s", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rData[gsId] = gsName
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("read kzinfo data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rData, nil
}

//获取券种 map[int]string ID-NAME
func (r *repZb) GetQzInfo() (map[int]string, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetQzInfo)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var gsId int
	var gsName string
	rData := make(map[int]string)
	for rows.Next() {
		err = rows.Scan(&gsId, &gsName)
		if err != nil {
			errMsg := fmt.Sprintf("read qzinfo data err: %s", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rData[gsId] = gsName
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("read qzinfo data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rData, nil
}

//获取数据（现金、赊账、交易次数）
func (r *repZb) GetBaoZhShouRSummaryData(mdId int, begDate time.Time, endDate time.Time) ([]*object.BaoZhShouRSummaryData, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetBaoZhShouRSummaryData,
		goToolCommon.GetDateStr(begDate), goToolCommon.GetDateStr(endDate), mdId)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rData := make([]*object.BaoZhShouRSummaryData, 0)
	for rows.Next() {
		d := object.BaoZhShouRSummaryData{}
		err = rows.Scan(&d.Hsr, &d.XjSr, &d.XjRate, &d.SzSr, &d.SzRate, &d.JyCs, &d.JyCsRate)
		if err != nil {
			errMsg := fmt.Sprintf("read BaoZhShouRSummaryData data err: %s", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rData = append(rData, &d)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("read BaoZhShouRSummaryData data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rData, nil
}

//获取转账明细数据
func (r *repZb) GetBaoZhShouRZzDetailData(mdId int, begDate time.Time, endDate time.Time) ([]*object.BaoZhShouRZzDetailData, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetBaoZhShouRZzDetailData,
		goToolCommon.GetDateStr(begDate), goToolCommon.GetDateStr(endDate), mdId)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rData := make([]*object.BaoZhShouRZzDetailData, 0)
	for rows.Next() {
		d := object.BaoZhShouRZzDetailData{}
		err = rows.Scan(&d.Hsr, &d.ZzJe, &d.ZzId)
		if err != nil {
			errMsg := fmt.Sprintf("read BaoZhShouRZzDetailData data err: %s", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rData = append(rData, &d)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("read BaoZhShouRZzDetailData data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rData, nil
}

//获取卡种明细数据
func (r *repZb) GetBaoZhShouRKzDetailData(mdId int, begDate time.Time, endDate time.Time) ([]*object.BaoZhShouRKzDetailData, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetBaoZhShouRKzDetailData,
		goToolCommon.GetDateStr(begDate), goToolCommon.GetDateStr(endDate), mdId)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rData := make([]*object.BaoZhShouRKzDetailData, 0)
	for rows.Next() {
		d := object.BaoZhShouRKzDetailData{}
		err = rows.Scan(&d.Hsr, &d.KzJe, &d.KzId)
		if err != nil {
			errMsg := fmt.Sprintf("read BaoZhShouRKzDetailData data err: %s", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rData = append(rData, &d)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("read BaoZhShouRKzDetailData data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rData, nil
}

//获取转账明细数据
func (r *repZb) GetBaoZhShouRQzDetailData(mdId int, begDate time.Time, endDate time.Time) ([]*object.BaoZhShouRQzDetailData, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetBaoZhShouRQzDetailData,
		goToolCommon.GetDateStr(begDate),
		goToolCommon.GetDateStr(endDate),
		mdId)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rData := make([]*object.BaoZhShouRQzDetailData, 0)
	for rows.Next() {
		d := object.BaoZhShouRQzDetailData{}
		err := rows.Scan(&d.Hsr, &d.QzJe, &d.QzId)
		if err != nil {
			errMsg := fmt.Sprintf("read BaoZhShouRZzDetailData data err: %s", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rData = append(rData, &d)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("read BaoZhShouRZzDetailData data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rData, nil
}

//获取可用的转账、券种、卡种
func (r *repZb) GetDisableList(mdId int) ([]string, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetDisableList, mdId)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rData := make([]string, 0)
	for rows.Next() {
		var n []byte
		err := rows.Scan(&n)
		if err != nil {
			errMsg := fmt.Sprintf("read DisableList data err: %s", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		n, err = goToolCommon.DecodeGB18030(n)
		if err != nil {
			errMsg := fmt.Sprintf("read DisableList data convert chinese err: %s", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rData = append(rData, string(n))
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("read DisableList data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rData, nil
}
