package log

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
}

func WriteLog(format string, v ...interface{}) {
	content := fmt.Sprintf(format, v...)
	logger.Println(content)
}
