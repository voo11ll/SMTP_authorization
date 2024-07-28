package logger

import (
	"auth/auth_back/pkg/utils"
	"os"
	"time"
)

func (l *Logger) openLogFile() (*os.File, error) {
	fileName := utils.ReturnDateString(time.Now(), "", false)
	fileOut, err := os.OpenFile("./logs/services/"+fileName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return fileOut, err
	}

	return fileOut, nil
}

// LogNotify is write information logs
func (l *Logger) LogNotify(message string, module string) {

	defer func() {
		if err := recover(); err != nil {
			l.LogError(err.(string), "logger.func.LogNotify")
		}
	}()

	fileOut, err := l.openLogFile()

	if err != nil {
		panic(err)
	}

	defer fileOut.Close()

	fileOut.WriteString("[INFO] " + utils.ReturnDateString(time.Now(), "/", true) + ": " + message + " (" + module + ")\n")
}

// LogWarning is write warning logs
func (l *Logger) LogWarning(message string, module string) {

	defer func() {
		if err := recover(); err != nil {
			l.LogError(err.(string), "logger.func.LogWarning")
		}
	}()

	fileOut, err := l.openLogFile()

	if err != nil {
		panic(err)
	}

	defer fileOut.Close()

	fileOut.WriteString("[WARNING] " + utils.ReturnDateString(time.Now(), "/", true) + ": " + message + " (" + module + ")\n")
}

// LogError is write error logs
func (l *Logger) LogError(message string, module string) {

	defer func() {
		if err := recover(); err != nil {
			l.LogError(err.(string), "logger.func.LogError")
		}
	}()

	fileOut, err := l.openLogFile()

	if err != nil {
		panic(err)
	}

	defer fileOut.Close()

	fileOut.WriteString("[ERROR] " + utils.ReturnDateString(time.Now(), "/", true) + ": " + message + " (" + module + ")\n")
}
