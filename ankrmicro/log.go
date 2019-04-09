package ankrmicro

import (
"fmt"
"log"
"path"
"runtime"
"time"
)


func WriteLog(msg string) {

	pathInfo := ""
	if _, file, line, ok := runtime.Caller(1); ok {

		pathInfo = fmt.Sprintf("[%s:%v]", path.Base(file), line)
	}

	time := time.Now().UTC().Format("2006-01-02 15:04:05")
	log.SetFlags(0)
	log.Println(time + "    " + msg + "    " + pathInfo )
}

