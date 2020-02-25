package logger

import (
	`log`
)

func LogErrors(err error){
	log.Println(err)
}
