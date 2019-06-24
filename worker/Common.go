package worker

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolSecret"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/global"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/object"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/repository"
	"math"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

type common struct {
}

func NewCommon() *common {
	return &common{}
}

//获取Token失效时长
func (c *common) GetTokenTimeout() time.Duration {
	return time.Duration(1000 * 1000 * 1000 * 60 * global.SysConfig.Web.TokenTimeout)
}

//获取Token
func (c *common) GetToken(mdId int, exp time.Duration) (string, error) {
	token := object.Token{
		MdId:   mdId,
		Expire: time.Now().Add(exp),
	}
	bToken, err := json.Marshal(token)
	if err != nil {
		errMsg := fmt.Sprintf("convert token to byte err: %s", err.Error())
		log.Error(errMsg)
		return "", errors.New(errMsg)
	}
	sToken, err := goToolSecret.EncryptToBase64Format(string(bToken), global.SecretKey)
	if err != nil {
		errMsg := fmt.Sprintf("encrypt token err: %s", err.Error())
		log.Error(errMsg)
		return "", errors.New(errMsg)
	}
	return sToken, nil
}

/*
验证Token
int token中的门店ID
bool token是否超时
error 错误
*/
func (c *common) VerifyToken(sToken string) (int, bool, error) {
	sToken, err := goToolSecret.DecryptFromBase64Format(sToken, global.SecretKey)
	if err != nil {
		errMsg := fmt.Sprintf("decrypt token err: %s", err.Error())
		log.Error(errMsg)
		return -1, false, errors.New(errMsg)
	}
	var token object.Token
	err = json.Unmarshal([]byte(sToken), &token)
	if err != nil {
		errMsg := fmt.Sprintf("convert byte to token err: %s", err.Error())
		log.Error(errMsg)
		return -1, false, errors.New(errMsg)
	}
	if !token.Expire.After(time.Now()) {
		errMsg := fmt.Sprintf("token is invalid")
		log.Error(errMsg)
		return -1, false, nil
	}
	return token.MdId, true, nil
}

//func (c *common) RefreshToken(sToken string)(string,error){
//	mdId,isVerified,err := c.VerifyToken(sToken)
//	if err != nil {
//		return "",err
//	}
//	if !isVerified {
//		return "",errors.New(string(object.ErrTypeMsgTokenTimeout))
//	}
//	return c.GetToken(mdId,c.GetTokenTimeout())
//}

//获取门店报账收入数据
func (c *common) GetMdBaoZhShouRData(mdId int, begDate time.Time, endDate time.Time) (
	zzStrList []string, kzStrList []string, qzStrList []string,
	data []*object.MdBaoZhShouRData, err error) {
	list := make(map[string]*object.MdBaoZhShouRData)
	for _, yyr := range c.getYyrList(begDate, endDate) {
		list[yyr] = &object.MdBaoZhShouRData{
			Yyr:            yyr,
			TransferDetail: make(map[string]float64),
			CardDetail:     make(map[string]float64),
			TicketDetail:   make(map[string]float64),
		}
	}
	rep := repository.NewRepZb()
	zzList, err := rep.GetZzInfo()
	if err != nil {
		return
	}
	for _, v := range zzList {
		zzStrList = append(zzStrList, v)
	}
	kzList, err := rep.GetKzInfo()
	if err != nil {
		return
	}
	for _, v := range kzList {
		kzStrList = append(kzStrList, v)
	}
	qzList, err := rep.GetQzInfo()
	if err != nil {
		return
	}
	for _, v := range qzList {
		qzStrList = append(qzStrList, v)
	}

	summaryData, err := rep.GetBaoZhShouRSummaryData(mdId, begDate, endDate)
	if err != nil {
		return
	}
	for _, sData := range summaryData {
		d, ok := list[goToolCommon.GetDateStr(sData.Hsr)]
		if !ok {
			continue
		}
		d.Cash = sData.XjSr * sData.XjRate
		d.Total = d.Total + d.Cash
		d.Credit = sData.SzSr * sData.SzRate
		d.Total = d.Total + d.Credit
		d.TotalCheck = int(math.Ceil(sData.JyCs * sData.JyCsRate))
	}

	zzData, err := rep.GetBaoZhShouRZzDetailData(mdId, begDate, endDate)
	if err != nil {
		return
	}
	for _, data := range zzData {
		d, ok := list[goToolCommon.GetDateStr(data.Hsr)]
		if !ok {
			continue
		}
		d.Total = d.Total + data.ZzJe
		d.Transfer = d.Transfer + data.ZzJe
		var zzName string
		zzName, ok = zzList[data.ZzId]
		if !ok {
			zzName = global.IsForbiddenTilte
		}
		_, ok = d.TransferDetail[zzName]
		if ok {
			d.TransferDetail[zzName] = d.TransferDetail[zzName] + data.ZzJe
		} else {
			d.TransferDetail[zzName] = data.ZzJe
		}
	}

	kzData, err := rep.GetBaoZhShouRKzDetailData(mdId, begDate, endDate)
	if err != nil {
		return
	}
	for _, data := range kzData {
		d, ok := list[goToolCommon.GetDateStr(data.Hsr)]
		if !ok {
			continue
		}
		d.Total = d.Total + data.KzJe
		d.Card = d.Card + data.KzJe
		var kzName string
		kzName, ok = kzList[data.KzId]
		if !ok {
			kzName = global.IsForbiddenTilte
		}
		_, ok = d.CardDetail[kzName]
		if ok {
			d.CardDetail[kzName] = d.CardDetail[kzName] + data.KzJe
		} else {
			d.CardDetail[kzName] = data.KzJe
		}
	}

	qzData, err := rep.GetBaoZhShouRQzDetailData(mdId, begDate, endDate)
	if err != nil {
		return
	}
	for _, data := range qzData {
		d, ok := list[goToolCommon.GetDateStr(data.Hsr)]
		if !ok {
			continue
		}
		d.Total = d.Total + data.QzJe
		d.Ticket = d.Ticket + data.QzJe
		var qzName string
		qzName, ok = qzList[data.QzId]
		if !ok {
			qzName = global.IsForbiddenTilte
		}
		_, ok = d.TicketDetail[qzName]
		if ok {
			d.TicketDetail[qzName] = d.TicketDetail[qzName] + data.QzJe
		} else {
			d.TicketDetail[qzName] = data.QzJe
		}
	}

	for _, v := range list {
		data = append(data, v)
	}
	return
}

//获取查询日期段内的营业日列表
func (c *common) getYyrList(begDate time.Time, endDate time.Time) []string {
	rList := make([]string, 0)
	for endDate.Add(time.Hour * 24).After(begDate) {
		rList = append(rList, goToolCommon.GetDateStr(begDate))
		begDate = begDate.Add(time.Hour * 24)
	}
	return rList
}
