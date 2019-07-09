package router

import (
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/global"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/object"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/repository"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/worker"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"sort"
	"strconv"
	"time"
)

type base struct {
	app *iris.Application
	c   common
}

func NewRouterBase(app *iris.Application) *base {
	return &base{
		app: app,
		c:   common{},
	}
}

func (base *base) AddBase() {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})
	v := base.app.Party("/", crs).AllowMethods(iris.MethodOptions)
	{
		v.Get("/version", base.version)
		v.Post("/login", base.login)
		v.Post("/logout", base.logout)
		v.Post("/refreshtoken", base.refreshToken)
		v.Post("/data", base.getMdData)
	}
}

//获取版本
func (base *base) version(ctx iris.Context) {
	v := object.VersionResponse{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
		Version: global.Version,
	}
	base.c.WriteResponse(ctx, v)
}

//刷新Token
func (base *base) refreshToken(ctx iris.Context) {
	var request object.RefreshTokenRequest
	var response object.RefreshTokenResponse
	err := ctx.ReadJSON(&request)
	if err != nil {
		response = object.RefreshTokenResponse{
			ErrCode: iris.StatusBadRequest,
			ErrMsg:  err.Error(),
			Token:   "",
		}
		base.c.WriteResponse(ctx, response)
		return
	}
	w := worker.NewCommon()
	mdId, isVerified, err := w.VerifyToken(request.Token)
	if err != nil {
		response = object.RefreshTokenResponse{
			ErrCode: -1,
			ErrMsg:  err.Error(),
			Token:   "",
		}
		base.c.WriteResponse(ctx, response)
		return
	}

	if !isVerified {
		response = object.RefreshTokenResponse{
			ErrCode: int(object.ErrTypeCodeTokenTimeout),
			ErrMsg:  string(object.ErrTypeMsgTokenTimeout),
			Token:   "",
		}
		base.c.WriteResponse(ctx, response)
		return
	}

	nt, err := w.GetToken(mdId, w.GetTokenTimeout())
	if err != nil {
		response = object.RefreshTokenResponse{
			ErrCode: -1,
			ErrMsg:  err.Error(),
			Token:   "",
		}
		base.c.WriteResponse(ctx, response)
		return
	}

	response = object.RefreshTokenResponse{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
		Token:   nt,
	}
	base.c.WriteResponse(ctx, response)
	return
}

//登录
func (base *base) login(ctx iris.Context) {
	var request object.LoginRequest
	var response object.LoginResponse
	err := ctx.ReadJSON(&request)
	if err != nil {
		response = object.LoginResponse{
			ErrCode: iris.StatusBadRequest,
			ErrMsg:  err.Error(),
		}
		base.c.WriteResponse(ctx, response)
		return
	}

	r := repository.NewRepZb()
	isVerified, err := r.LoginVerify(request.LoginName, goToolCommon.Md5([]byte(request.LoginPwd)))
	if err != nil {
		response = object.LoginResponse{
			ErrCode: -1,
			ErrMsg:  err.Error(),
		}
		base.c.WriteResponse(ctx, response)
		return
	}

	if !isVerified {
		response = object.LoginResponse{
			ErrCode: int(object.ErrTypeCodeLoginFailed),
			ErrMsg:  string(object.ErrTypeMsgLoginFailed),
		}
		base.c.WriteResponse(ctx, response)
		return
	}

	mdId, err := r.GetMdIdByLogin(request.LoginName)
	if err != nil {
		response = object.LoginResponse{
			ErrCode: -1,
			ErrMsg:  fmt.Sprintf("login success,but get mdid err: %s", err.Error()),
		}
		base.c.WriteResponse(ctx, response)
		return
	}

	mdName, err := r.GetMdName(mdId)
	if err != nil {
		response = object.LoginResponse{
			ErrCode: -1,
			ErrMsg:  fmt.Sprintf("login success,but get mdname err: %s", err.Error()),
		}
		base.c.WriteResponse(ctx, response)
		return
	}

	w := worker.NewCommon()
	token, err := w.GetToken(mdId, w.GetTokenTimeout())
	if err != nil {
		response = object.LoginResponse{
			ErrCode: -1,
			ErrMsg:  err.Error(),
		}
		base.c.WriteResponse(ctx, response)
		return
	}

	response = object.LoginResponse{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
		MdName:  mdName,
		Token:   token,
	}
	base.c.WriteResponse(ctx, response)
	return
}

//登出
func (base *base) logout(ctx iris.Context) {
	v := object.Response{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
	}
	base.c.WriteResponse(ctx, v)
}

//获取门店数据
func (base *base) getMdData(ctx iris.Context) {
	var request object.GetMdDataRequest
	var response object.GetMdDataResponse
	err := ctx.ReadJSON(&request)
	if err != nil {
		response = object.GetMdDataResponse{
			ErrCode: iris.StatusBadRequest,
			ErrMsg:  err.Error(),
		}
		base.c.WriteResponse(ctx, response)
		return
	}
	w := worker.NewCommon()
	mdId, isVerified, err := w.VerifyToken(request.Token)
	if err != nil {
		response = object.GetMdDataResponse{
			ErrCode: -1,
			ErrMsg:  err.Error(),
		}
		base.c.WriteResponse(ctx, response)
		return
	}
	if !isVerified {
		response = object.GetMdDataResponse{
			ErrCode: int(object.ErrTypeCodeTokenTimeout),
			ErrMsg:  string(object.ErrTypeMsgTokenTimeout),
		}
		base.c.WriteResponse(ctx, response)
		return
	}
	sTime, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		errMsg := fmt.Sprintf("convert start date err: %s", err.Error())
		response = object.GetMdDataResponse{
			ErrCode: -1,
			ErrMsg:  errMsg,
		}
		base.c.WriteResponse(ctx, response)
		return
	}
	eTime, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		errMsg := fmt.Sprintf("convert end date err: %s", err.Error())
		response = object.GetMdDataResponse{
			ErrCode: -1,
			ErrMsg:  errMsg,
		}
		base.c.WriteResponse(ctx, response)
		return
	}
	//开始日期不能大于截止日期，否则互换
	if sTime.After(eTime) {
		t := sTime
		sTime = eTime
		eTime = t
	}
	//日期间隔不能大于XX天
	if sTime.Add(goToolCommon.GetDurationByDay(global.SysConfig.Web.MaxSearchDays)).Before(eTime.Add(time.Second)) {
		errMsg := fmt.Sprintf(string(object.ErrTypeMsgExceedMaxSearchDay), global.SysConfig.Web.MaxSearchDays)
		response = object.GetMdDataResponse{
			ErrCode: int(object.ErrTypeCodeExceedMaxSearchDay),
			ErrMsg:  errMsg,
		}
		base.c.WriteResponse(ctx, response)
		return
	}

	zzList, kzList, qzList, d, err := w.GetMdBaoZhShouRData(mdId, sTime, eTime)
	if err != nil {
		response = object.GetMdDataResponse{
			ErrCode: -1,
			ErrMsg:  err.Error(),
		}
		base.c.WriteResponse(ctx, response)
		return
	}

	total := object.MdBaoZhShouRData{
		Yyr:            "合计",
		Total:          0,
		Cash:           0,
		Credit:         0,
		Transfer:       0,
		TransferDetail: make(map[string]float64),
		Card:           0,
		CardDetail:     make(map[string]float64),
		Ticket:         0,
		TicketDetail:   make(map[string]float64),
		TotalCheck:     0,
	}
	rKey := make([]string, 0)
	rList := make(map[string]object.GetMdDataResponseDetail)
	for _, dd := range d {
		total.Total = total.Total + dd.Total
		total.Cash = total.Cash + dd.Cash
		total.Credit = total.Credit + dd.Credit
		total.Transfer = total.Transfer + dd.Transfer
		total.Card = total.Card + dd.Card
		total.Ticket = total.Ticket + dd.Ticket
		total.TotalCheck = total.TotalCheck + dd.TotalCheck
		zd := make(map[string]string)
		for _, n := range zzList {
			_, ok := dd.TransferDetail[n]
			if ok {
				zd["transfer"+n] = strconv.FormatFloat(dd.TransferDetail[n], 'f', 2, 64)
			} else {
				zd["transfer"+n] = ""
			}
			_, ok = total.TransferDetail[n]
			if ok {
				total.TransferDetail[n] = total.TransferDetail[n] + dd.TransferDetail[n]
			} else {
				total.TransferDetail[n] = dd.TransferDetail[n]
			}
		}
		kd := make(map[string]string)
		for _, n := range kzList {
			_, ok := dd.CardDetail[n]
			if ok {
				kd["card"+n] = strconv.FormatFloat(dd.CardDetail[n], 'f', 2, 64)
			} else {
				kd["card"+n] = ""
			}
			_, ok = total.CardDetail[n]
			if ok {
				total.CardDetail[n] = total.CardDetail[n] + dd.CardDetail[n]
			} else {
				total.CardDetail[n] = dd.CardDetail[n]
			}
		}
		qz := make(map[string]string)
		for _, n := range qzList {
			_, ok := dd.TicketDetail[n]
			if ok {
				qz["ticket"+n] = strconv.FormatFloat(dd.TicketDetail[n], 'f', 2, 64)
			} else {
				qz["ticket"+n] = ""
			}
			_, ok = total.TicketDetail[n]
			if ok {
				total.TicketDetail[n] = total.TicketDetail[n] + dd.TicketDetail[n]
			} else {
				total.TicketDetail[n] = dd.TicketDetail[n]
			}
		}

		var zf, kf, qf string
		_, zzOk := dd.TransferDetail[global.IsForbiddenTitle]
		if zzOk {
			zf = strconv.FormatFloat(dd.TransferDetail[global.IsForbiddenTitle], 'f', 2, 64)
			_, ok := total.TransferDetail[global.IsForbiddenTitle]
			if ok {
				total.TransferDetail[global.IsForbiddenTitle] = total.TransferDetail[global.IsForbiddenTitle] + dd.TransferDetail[global.IsForbiddenTitle]
			} else {
				total.TransferDetail[global.IsForbiddenTitle] = dd.TransferDetail[global.IsForbiddenTitle]
			}
		} else {
			zf = ""
		}
		_, kzOk := dd.CardDetail[global.IsForbiddenTitle]
		if kzOk {
			kf = strconv.FormatFloat(dd.CardDetail[global.IsForbiddenTitle], 'f', 2, 64)
			_, ok := total.CardDetail[global.IsForbiddenTitle]
			if ok {
				total.CardDetail[global.IsForbiddenTitle] = total.CardDetail[global.IsForbiddenTitle] + dd.CardDetail[global.IsForbiddenTitle]
			} else {
				total.CardDetail[global.IsForbiddenTitle] = dd.CardDetail[global.IsForbiddenTitle]
			}
		} else {
			kf = ""
		}
		_, qzOk := dd.TicketDetail[global.IsForbiddenTitle]
		if qzOk {
			qf = strconv.FormatFloat(dd.TicketDetail[global.IsForbiddenTitle], 'f', 2, 64)
			_, ok := total.TicketDetail[global.IsForbiddenTitle]
			if ok {
				total.TicketDetail[global.IsForbiddenTitle] = total.TicketDetail[global.IsForbiddenTitle] + dd.TicketDetail[global.IsForbiddenTitle]
			} else {
				total.TicketDetail[global.IsForbiddenTitle] = dd.TicketDetail[global.IsForbiddenTitle]
			}
		} else {
			qf = ""
		}

		rKey = append(rKey, dd.Yyr)
		rList[dd.Yyr] = object.GetMdDataResponseDetail{
			Yyr:               dd.Yyr,
			Total:             strconv.FormatFloat(dd.Total, 'f', 2, 64),
			Cash:              strconv.FormatFloat(dd.Cash, 'f', 2, 64),
			Credit:            strconv.FormatFloat(dd.Credit, 'f', 2, 64),
			Transfer:          strconv.FormatFloat(dd.Transfer, 'f', 2, 64),
			TransferDetail:    zd,
			TransferForbidden: zf,
			Card:              strconv.FormatFloat(dd.Card, 'f', 2, 64),
			CardDetail:        kd,
			CardForbidden:     kf,
			Ticket:            strconv.FormatFloat(dd.Ticket, 'f', 2, 64),
			TicketDetail:      qz,
			TicketForbidden:   qf,
			TotalCheck:        strconv.Itoa(dd.TotalCheck),
		}
	}
	rKey = append(rKey, total.Yyr)

	tzd := make(map[string]string)
	for _, n := range zzList {
		_, ok := total.TransferDetail[n]
		if ok {
			tzd["transfer"+n] = strconv.FormatFloat(total.TransferDetail[n], 'f', 2, 64)
		} else {
			tzd["transfer"+n] = ""
		}
	}
	tkd := make(map[string]string)
	for _, n := range kzList {
		_, ok := total.CardDetail[n]
		if ok {
			tkd["card"+n] = strconv.FormatFloat(total.CardDetail[n], 'f', 2, 64)
		} else {
			tkd["card"+n] = ""
		}
	}
	tqz := make(map[string]string)
	for _, n := range qzList {
		_, ok := total.TicketDetail[n]
		if ok {
			tqz["ticket"+n] = strconv.FormatFloat(total.TicketDetail[n], 'f', 2, 64)
		} else {
			tqz["ticket"+n] = ""
		}
	}

	var tzf, tkf, tqf string
	_, zzOk := total.TransferDetail[global.IsForbiddenTitle]
	if zzOk {
		tzf = strconv.FormatFloat(total.TransferDetail[global.IsForbiddenTitle], 'f', 2, 64)
	} else {
		tzf = ""
	}
	_, kzOk := total.CardDetail[global.IsForbiddenTitle]
	if kzOk {
		tkf = strconv.FormatFloat(total.CardDetail[global.IsForbiddenTitle], 'f', 2, 64)
	} else {
		tkf = ""
	}
	_, qzOk := total.TicketDetail[global.IsForbiddenTitle]
	if qzOk {
		tqf = strconv.FormatFloat(total.TicketDetail[global.IsForbiddenTitle], 'f', 2, 64)
	} else {
		tqf = ""
	}

	rTotal := object.GetMdDataResponseDetail{
		Yyr:               total.Yyr,
		Total:             strconv.FormatFloat(total.Total, 'f', 2, 64),
		Cash:              strconv.FormatFloat(total.Cash, 'f', 2, 64),
		Credit:            strconv.FormatFloat(total.Credit, 'f', 2, 64),
		Transfer:          strconv.FormatFloat(total.Transfer, 'f', 2, 64),
		TransferDetail:    tzd,
		TransferForbidden: tzf,
		Card:              strconv.FormatFloat(total.Card, 'f', 2, 64),
		CardDetail:        tkd,
		CardForbidden:     tkf,
		Ticket:            strconv.FormatFloat(total.Ticket, 'f', 2, 64),
		TicketDetail:      tqz,
		TicketForbidden:   tqf,
		TotalCheck:        strconv.Itoa(total.TotalCheck),
	}

	rList[total.Yyr] = rTotal

	//sort.Strings(rKey)
	sort.Sort(goToolCommon.SortByPinyin(rKey))
	rRList := make([]object.GetMdDataResponseDetail, 0)
	for i := 0; i < len(rList); i++ {
		rRList = append(rRList, rList[rKey[i]])
	}

	//sort.Strings(zzList)
	sort.Sort(goToolCommon.SortByPinyin(zzList))
	//sort.Strings(kzList)
	sort.Sort(goToolCommon.SortByPinyin(kzList))
	//sort.Strings(qzList)
	sort.Sort(goToolCommon.SortByPinyin(qzList))

	response = object.GetMdDataResponse{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
		ZzList:  zzList,
		KzList:  kzList,
		QzList:  qzList,
		Data:    rRList,
		//Total:   strTotal,
	}
	base.c.WriteResponse(ctx, response)
	return
}
