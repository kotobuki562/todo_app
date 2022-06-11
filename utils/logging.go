package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSetting(logFile string) {
	// Logファイルを作成するか書込みを行う
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	// ファイルの書き込み
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

/*
main.goで実行するとログが作成される
*/