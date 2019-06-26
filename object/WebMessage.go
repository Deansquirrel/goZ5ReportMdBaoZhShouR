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
	MdName  string `json:"mdname"`
	Token   string `json:"token"`
}

type GetMdDataRequest struct {
	Token     string `json:"token"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}

type GetMdDataResponse struct {
	ErrCode int                       `json:"errcode"`
	ErrMsg  string                    `json:"errmsg"`
	ZzList  []string                  `json:"zzlist"`
	KzList  []string                  `json:"kzlist"`
	QzList  []string                  `json:"qzlist"`
	Data    []GetMdDataResponseDetail `json:"data"`
}

type GetMdDataResponseDetail struct {
	Yyr               string            `json:"yyr"`
	Total             string            `json:"total"`             //合计
	Cash              string            `json:"cash"`              //现金
	Credit            string            `json:"credit"`            //赊账
	Transfer          string            `json:"transfer"`          //转账
	TransferDetail    map[string]string `json:"transferdetail"`    //转账明细
	TransferForbidden string            `json:"transferforbidden"` //转账已禁用`
	Card              string            `json:"card"`              //卡种
	CardDetail        map[string]string `json:"carddetail"`        //卡种明细
	CardForbidden     string            `json:"cardforbidden"`     //卡种已禁用`
	Ticket            string            `json:"ticket"`            //券种
	TicketDetail      map[string]string `json:"ticketdetail"`      //券种明细
	TicketForbidden   string            `json:"ticketforbidden"`   //券种已禁用
	TotalCheck        string            `json:"totalcheck"`        //交易次数
}
