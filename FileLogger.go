package FileLogger

import (
	"os"
	"time"
)

type FileLogger struct {
	currentFile *os.File
}

func NewLogger() *FileLogger {
	return &FileLogger{
		currentFile: nil,
	}
}

func getDate(now time.Time) string {
	return now.Format("02-01-06")
}

func getTime(now time.Time) string {
	return now.Format("15:04:05")
}

func makeFilename() string {
	now := time.Now()
	return "log_" + getDate(now) + "." + getTime(now) + ".txt"
}

func (fl *FileLogger) Log(logMessage string) {
	if fl.currentFile == nil {
		newLogFile, err := os.Create(makeFilename())
		if err != nil {
			panic(err)
		}
		fl.currentFile = newLogFile
	} else {
		fileReopened, err := os.OpenFile(fl.currentFile.Name(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		fl.currentFile = fileReopened
	}
	now := time.Now()
	_, err := fl.currentFile.WriteString("At: " + getTime(now) + " " + logMessage + "\n")
	if err != nil {
		panic(err)
	}
	err = fl.currentFile.Close()
	if err != nil {
		panic(err)
	}
}
