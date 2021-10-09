package logging

import (
	"fmt"
	"gin-api/pkg/setting"
	"log"
	"os"
	"time"
)

var (
	LogPath = setting.AppSetting.LogSavePath
	LogFilePrefix = setting.AppSetting.LogSaveName
	LogFileExt = setting.AppSetting.LogFileExt
	TimeFormat = setting.AppSetting.TimeFormat
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
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