package logger

import (
	"fmt"
	"log"

	"github.com/vinay03/chalk"
)

func Info(v ...interface{}) {
	msg := fmt.Sprint(v...)
	log.Println(chalk.Black().BgBlue(" ℹ️  INFO: " + msg))
}

func Warning(v ...interface{}) {
	msg := fmt.Sprint(v...)
	log.Println(chalk.Black().BgYellow(" ⚠️ WARNING: " + msg))
}

func Success(v ...interface{}) {
	msg := fmt.Sprint(v...)
	log.Println(chalk.Black().BgGreen(" ✅ SUCCESS: " + msg))
}

func Error(v ...interface{}) {
	msg := fmt.Sprint(v...)
	log.Println(chalk.Black().BgRed(" ❌ ERROR: " + msg))
}

func Fatal(v ...interface{}) {
	msg := fmt.Sprint(v...)
	log.Fatal(chalk.Black().BgRed(" 💀 FATAL: " + msg))
}
