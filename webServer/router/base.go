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
	var response object.GetMdDatResponse
	err := ctx.ReadJSON(&request)
	if err != nil {
		response = object.GetMdDatResponse{
			ErrCode: iris.StatusBadRequest,
			ErrMsg:  err.Error(),
		}
		base.c.WriteResponse(ctx, response)
		return
	}
	w := worker.NewCommon()
	mdId, isVerified, err := w.VerifyToken(request.Token)
	if err != nil {
		response = object.GetMdDatResponse{
			ErrCode: -1,
			ErrMsg:  err.Error(),
		}
		base.c.WriteResponse(ctx, response)
		return
	}
	if !isVerified {
		response = object.GetMdDatResponse{
			ErrCode: int(object.ErrTypeCodeTokenTimeout),
			ErrMsg:  string(object.ErrTypeMsgTokenTimeout),
		}
		base.c.WriteResponse(ctx, response)
		return
	}
	sTime, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		errMsg := fmt.Sprintf("convert start date err: %s", err.Error())
		response = object.GetMdDatResponse{
			ErrCode: -1,
			ErrMsg:  errMsg,
		}
		base.c.WriteResponse(ctx, response)
		return
	}
	eTime, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		errMsg := fmt.Sprintf("convert end date err: %s", err.Error())
		response = object.GetMdDatResponse{
			ErrCode: -1,
			ErrMsg:  errMsg,
		}
		base.c.WriteResponse(ctx, response)
		return
	}
	d, err := w.GetMdBaoZhShouRData(mdId, sTime, eTime)
	if err != nil {
		response = object.GetMdDatResponse{
			ErrCode: -1,
			ErrMsg:  err.Error(),
		}
		base.c.WriteResponse(ctx, response)
		return
	}

	response = object.GetMdDatResponse{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
		Data:    d,
	}
	base.c.WriteResponse(ctx, response)
	return
}
