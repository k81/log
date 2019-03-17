package log

type Level uint8

type LevelName string

const (
	LevelTrace Level = 1 << iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
	LevelMask = 0xFF
)

func (level Level) String() string {
	levelName := "UNKNOWN"
	switch level {
	case LevelTrace:
		levelName = "TRACE"
	case LevelDebug:
		levelName = "DEBUG"
	case LevelInfo:
		levelName = "INFO"
	case LevelWarning:
		levelName = "WARNING"
	case LevelError:
		levelName = "ERROR"
	case LevelFatal:
		levelName = "FATAL"
	}
	return levelName
}

func (levelName LevelName) ToLevel() Level {
	level := LevelDebug
	switch levelName {
	case "TRACE":
		level = LevelTrace
	case "DEBUG":
		level = LevelDebug
	case "INFO":
		level = LevelInfo
	case "WARNING":
		level = LevelWarning
	case "ERROR":
		level = LevelError
	case "FATAL":
		level = LevelFatal
	}
	return level
}
