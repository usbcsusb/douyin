package setting

import (
	"go.uber.org/zap"
	"toktik-common/logger"
	"toktik-interaction/internal/global"
)

func init() {
	l := log{}
	Settings = append(Settings, l)
}

type log struct {
}

// InitSetting 日志初始化
func (log) InitSetting() {
	global.Logger = logger.NewLogger(&logger.InitStruct{
		LogSavePath:   global.Settings.Logger.LogSavePath,
		LogFileExt:    global.Settings.Logger.LogFileExt,
		MaxSize:       global.Settings.Logger.MaxSize,
		MaxBackups:    global.Settings.Logger.MaxBackups,
		MaxAge:        global.Settings.Logger.MaxAge,
		Compress:      global.Settings.Logger.Compress,
		LowLevelFile:  global.Settings.Logger.LowLevelFile,
		HighLevelFile: global.Settings.Logger.HighLevelFile,
	}, global.Settings.Logger.Level)
	zap.ReplaceGlobals(global.Logger.Logger)
}
