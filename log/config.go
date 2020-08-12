package log

type ConfigFileWriter struct {
	On              bool   `toml:"On"`
	LogPath         string `toml:"LogPath"`
	RotateLogPath   string `toml:"RotateLogPath"`
	WfLogPath       string `toml:"WfLogPath"`
	RotateWfLogPath string `toml:"RotateWfLogPath"`
}

type ConfConsoleWriter struct {
	On    bool `toml:"On"`
	Color bool `toml:"On"`
}

type Config struct {
	Level string            `toml:"LogLevel"`
	FW    ConfConsoleWriter `toml:"FileWriter"`
	CW    ConfConsoleWriter `toml:"ConsoleWriter"`
}
