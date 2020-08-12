package log

import (
	"testing"
	"time"
)

//测试日志实例打点
func TestLogInstance(t *testing.T) {
	nlog := NewLogger()
	logConf := Config{
		Level: "trace",
		FW: ConfigFileWriter{
			On:              true,
			LogPath:         "./log_test.log",
			RotateLogPath:   "./log_test.log",
			WfLogPath:       "./log_test.wf.log",
			RotateWfLogPath: "./log_test.wf.log",
		},
		CW: ConfConsoleWriter{
			On:    true,
			Color: true,
		},
	}
	SetupInstanceWithConf(logConf, nlog)
	nlog.Warn("test message")
	time.Sleep(11 * time.Second)
	nlog.Close()

}
