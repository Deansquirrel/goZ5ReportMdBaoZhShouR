package global

import (
	"context"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/object"
)

const (
	//PreVersion = "0.0.0 Build20190101"
	//TestVersion = "0.0.0 Build20190101"
	Version = "0.0.0 Build20190101"

	SecretKey        = "Z5ReportMdBaoZhShouR"
	IsForbiddenTilte = "已禁用"
)

var Ctx context.Context
var Cancel func()

//程序启动参数
var Args *object.ProgramArgs

//系统参数
var SysConfig *object.SystemConfig
