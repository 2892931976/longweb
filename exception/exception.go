package exception

import (
	"fmt"
	"os"
	"runtime"

	. "github.com/devfeel/longweb/const"
	"github.com/devfeel/longweb/framework/log"
)

//统一异常处理
func CatchError(title string, logtarget string, err interface{}) (errmsg string) {
	errmsg = fmt.Sprintln(err)
	os.Stdout.Write([]byte(title + " error! => " + errmsg + " => "))
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, true)
	logger.Log(title+" error! => "+errmsg+" => "+string(buf[:n]), logtarget, LogLevel_Error)
	return errmsg
}
