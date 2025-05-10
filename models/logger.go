package models

import "os"

var logFile, _ = os.Create("log.txt")

func WriteToLog(text string) {
	logFile.WriteString(text)
}

func CloseLogFile() {
	logFile.Close()
}
