package lib

import (
	"bytes"
	"github.com/julianlee107/go-common/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

type BaseConf struct {
	DebugMode    string    `mapstructure:"debug_mode"`
	TimeLocation string    `mapstructure:"time_location"`
	Log          LogConfig `mapstruacture:"log"`
	Base         struct {
		DebugMode    string `mapstructure:"debug_mode"`
		TimeLocation string `mapstructure:"time_location"`
	} `mapstructure:"base"`
}
type LogConfFileWriter struct {
	On              bool   `mapstructure:"on"`
	LogPath         string `mapstructure:"log_path"`
	RotateLogPath   string `mapstructure:"rotate_log_path"`
	WfLogPath       string `mapstructure:"wf_log_path"`
	RotateWfLogPath string `mapstructure:"rotate_wf_log_path"`
}

type LogConfConsoleWriter struct {
	On    bool `mapstructure:"on"`
	Color bool `mapstructure:"color"`
}

type LogConfig struct {
	Level string               `mapstructure:"log_level"`
	FW    LogConfFileWriter    `mapstructure:"file_writer"`
	CW    LogConfConsoleWriter `mapstructure:"console_writer"`
}

var (
	ConfBase     *BaseConf
	ViperConfMap map[string]*viper.Viper
)

// get base config
func GetBaseConf() *BaseConf {
	return ConfBase
}

func InitBaseConf(path string) error {
	ConfBase = &BaseConf{}
	err := ParseConfig(path, ConfBase)
	if err != nil {
		return err
	}
	if ConfBase.DebugMode == "" {
		if ConfBase.Base.DebugMode != "" {
			ConfBase.DebugMode = ConfBase.Base.DebugMode
		} else {
			ConfBase.DebugMode = "debug"
		}
	}

	if ConfBase.TimeLocation == "" {
		if ConfBase.Base.TimeLocation != "" {
			ConfBase.TimeLocation = ConfBase.Base.TimeLocation
		} else {
			ConfBase.TimeLocation = "Asia/Chongqing"
		}
	}

	logConf := log.Config{
		Level: ConfBase.Log.Level,
		FW: log.ConfigFileWriter{
			On:              ConfBase.Log.FW.On,
			LogPath:         ConfBase.Log.FW.LogPath,
			RotateLogPath:   ConfBase.Log.FW.RotateLogPath,
			RotateWfLogPath: ConfBase.Log.FW.RotateWfLogPath,
			WfLogPath:       ConfBase.Log.FW.WfLogPath,
		},
		CW: log.ConfConsoleWriter{
			Color: ConfBase.Log.CW.Color,
			On:    ConfBase.Log.CW.On,
		},
	}
	if err := log.SetupDefaultLogWithConf(logConf); err != nil {
		panic(err)
	}
	log.SetLayout("2006-01-02 15:04:05.000")
	return nil
}

func InitViperConf() error {
	file, err := os.Open(ConfEnvPath + string(os.PathSeparator))
	if err != nil {
		return err
	}
	fileList, err := file.Readdir(1024)
	if err != nil {
		return err
	}
	for _, f := range fileList {
		if !f.IsDir() {
			conf, err := ioutil.ReadFile(ConfEnvPath + string(os.PathSeparator) + f.Name())
			if err != nil {
				return err
			}
			v := viper.New()
			confType := strings.Split(f.Name(), ".")
			v.SetConfigType(confType[len(confType)-1])
			v.ReadConfig(bytes.NewBuffer(conf))
			if ViperConfMap == nil {
				ViperConfMap = make(map[string]*viper.Viper)
			}
			ViperConfMap[confType[0]] = v
		}
	}
	return nil
}
