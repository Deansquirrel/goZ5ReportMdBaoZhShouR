package object

import "time"

//钉钉消息发送配置
type DingTalkRobotConfigData struct {
	FWebHookKey string
	FAtMobiles  string
	FIsAtAll    int
}

type Token struct {
	MdId   int
	Expire time.Time
}

type MdBaoZhShouRData struct {
	Date           string             `json:"date"`
	Total          float64            `json:"total"`          //合计
	Cash           float64            `json:"cash"`           //现金
	Credit         float64            `json:"credit"`         //赊账
	Transfer       float64            `json:"transfer"`       //转账
	TransferDetail map[string]float64 `json:"transferdetail"` //转账明细
	Card           float64            `json:"card"`           //卡种
	CardDetail     map[string]float64 `json:"carddetail"`     //卡种明细
	Ticket         float64            `json:"ticket"`         //券种
	TicketDetail   map[string]float64 `json:"ticketdetail"`   //券种明细
	TotalCheck     int                `json:"totalcheck"`     //交易次数
}

type mdBaoZhShouRDetailData struct {
}
