package logger

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strconv"
)

func Log(err error) {
	if err == nil {
		msg := "you cannot pass a nil error to the logger.Log function"
		newErr := errors.New(msg)
		
		if newErr == nil {
			fmt.Println(msg)
			return
		}

		Log(newErr)
		return
	}

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
