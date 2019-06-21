package object

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type VersionResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Version string `json:"version"`
}

type RefreshTokenRequest struct {
	Token string `json:"token"`
}

type RefreshTokenResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Token   string `json:"token"`
}

type LoginRequest struct {
	LoginName string `json:"loginname"`
	LoginPwd  string `json:"loginpwd"`
}

type LoginResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Token   string `json:"token"`
}

type GetMdDataRequest struct {
	Token     string `json:"token"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}

type GetMdDatResponse struct {
	ErrCode int                `json:"errcode"`
	ErrMsg  string             `json:"errmsg"`
	Data    []MdBaoZhShouRData `json:"data"`
}
