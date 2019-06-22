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

type BaoZhShouRSummaryData struct {
	Hsr      time.Time
	XjSr     float64
	XjRate   float64
	SzSr     float64
	SzRate   float64
	JyCs     float64
	JyCsRate float64
}

type BaoZhShouRZzDetailData struct {
	Hsr  time.Time
	ZzJe float64
	ZzId int
}

type BaoZhShouRKzDetailData struct {
	Hsr  time.Time
	KzJe float64
	KzId int
}

type BaoZhShouRQzDetailData struct {
	Hsr  time.Time
	QzJe float64
	QzId int
}

type MdBaoZhShouRData struct {
	Yyr            string             `json:"yyr"`
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
