package logger

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// GenerateLogger 生成日志引擎
// dirPath：日志保存路径（文件夹），如果输入的是文件，则截取文件夹路径，支持. ~，格式为./log
// cycle：日志保存周期，即按天(day)，按月(month)，按年(year)等方式进行保存
func GenerateLogger(dirPath string, cycle string) (zerolog.Logger, error) {
	var filePath string

	dirPath = strings.TrimSpace(dirPath)

	if filepath.Ext(dirPath) != "" {
		return zerolog.Logger{}, errors.New("dirPath格式错误")
	}

	lastch := string([]rune(dirPath)[len([]rune(dirPath))-1:])
	if lastch != "/" {
		dirPath += "/"
	}

	ch := string([]rune(dirPath)[1:])
	firstch := string([]rune(dirPath)[:1])
	if firstch != "." && firstch != "~" && firstch != "/" {
		return zerolog.Logger{}, errors.New("dirPath格式错误")
	}

	switch firstch {
	case ".":
		pwd, _ := os.Getwd()
		dirPath = pwd + ch
	case "~":
		h, _ := homedir.Dir()
		dirPath = h + ch
	}

	dirPath = filepath.Dir(dirPath)

	switch strings.ToLower(cycle) {
	case "day":
		filePath = fmt.Sprintf("%s/%s.log", dirPath, time.Now().Format("20060102"))
	case "month":
		filePath = fmt.Sprintf("%s/%s.log", dirPath, time.Now().Format("200601"))
	case "year":
		filePath = fmt.Sprintf("%s/%s.log", dirPath, time.Now().Format("200601"))
	default:
		filePath = fmt.Sprintf("%s/log.log", dirPath)
	}

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDONLY, 0666)
	if err != nil {
		return zerolog.Logger{}, err
	}

	zerolog.TimeFieldFormat = "01-02 15:04:05"
	log.Logger = zerolog.New(f).With().Timestamp().Caller().Logger()

	return log.Logger, nil
}
