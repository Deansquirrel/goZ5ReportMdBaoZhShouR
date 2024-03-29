package global

import (
	"context"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/object"
)

const (
	//PreVersion = "1.0.6 Build20191218"
	//TestVersion = "0.0.0 Build20190101"
	Version = "1.0.7 Build20191219"

	SecretKey        = "Z5ReportMdBaoZhShouR"
	IsForbiddenTitle = "已禁用"
)

var Ctx context.Context
var Cancel func()

//程序启动参数
var Args *object.ProgramArgs

//系统参数
var SysConfig *object.SystemConfig
