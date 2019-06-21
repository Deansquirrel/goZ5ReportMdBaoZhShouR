package worker

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Deansquirrel/goToolSecret"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/global"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/object"
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
