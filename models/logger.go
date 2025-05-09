package models

import "os"

var logFile, _ = os.Create("log.txt")

func Log(text string) {
	logFile.WriteString(text)
}

func Close() {
	logFile.Close()
}
