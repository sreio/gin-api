package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogPath = "runtime/logs/"
	LogFilePrefix = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprint(LogPath)
}

func getLogFileFullPath() string {
	filepath := getLogFilePath()
	fileFullPath := fmt.Sprintf("%s%s.%s", LogFilePrefix, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", filepath, fileFullPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
		case os.IsNotExist(err) :
			mkDir()
		case os.IsPermission(err) :
			log.Fatalf("Permission :%v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to openFile Error:%s", err)
	}
	return handle
}

func mkDir() {
	dir,_ := os.Getwd()
	err := os.MkdirAll(dir + "/" + getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}