package logger

import (
	"io"
	"os"

	"website-conector/utils"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger
var IS_DEBUG = false

func InitLog() {
	utils.LoadVars()
	writer := io.Writer(os.Stdout)
	Log = zerolog.New(writer)
	debug := os.Getenv("DEBUG")
	switch debug {
	case "":
		IS_DEBUG = false
	case "true":
		IS_DEBUG = true
	case "false":
		IS_DEBUG = false
	}
}

func Error(from string, msg string) {
	Log.Log().Timestamp().Str("level", "error").Str("from", from).Str("message", msg).Send()
}

func Info(from string, msg string) {
	Log.Log().Timestamp().Str("level", "info").Str("from", from).Str("message", msg).Send()
}

func Fatal(from string, msg string) {
	Log.Log().Timestamp().Str("level", "fatal").Str("from", from).Str("message", msg).Send()
	os.Exit(1)
}

func Warn(from string, msg string) {
	Log.Log().Timestamp().Str("level", "warn").Str("from", from).Str("message", msg).Send()
}

func Panic(from string, msg string) {
	Log.Log().Timestamp().Str("level", "panic").Str("from", from).Str("message", msg).Send()
	panic(msg)
}

func Debug(from string, msg string) {
	if IS_DEBUG {
		Log.Log().Timestamp().Str("level", "debug").Str("from", from).Str("message", msg).Send()
	}
}
