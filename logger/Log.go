package logger

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
)

func Log(err error) {
	pc := make([]uintptr, 50)
	callers := runtime.Callers(1, pc)
	callStrs := ""
	for i := 1; i <= callers; i++ {
		_, file, line, _ := runtime.Caller(i)
		callStr := file + ": " + "line " + strconv.Itoa(line) + "\n"
		callStrs = callStrs + callStr
	}

	msg := fmt.Errorf("ERROR: %w\n"+callStrs, err)

	log.Println(msg)
}
