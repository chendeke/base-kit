package level

import (
	"github.com/chendeke/config/config/reader"
	"github.com/chendeke/logs"
)

const (
	LevelPath = "logs.level"
)

func ChangeLogLevel(value reader.Value) {
	level := value.String("")
	if len(level) > 0 && (level == "debug" || level == "info" || level == "error" ||
		level == "warn" || level == "panic" || level == "fatal") {
		logs.SetLevel(level)
	}
}
