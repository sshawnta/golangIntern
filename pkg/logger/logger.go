package logger

import (
	`log`
)

// LogErrors reports an error
func LogErrors(err error) {
	log.Println(err)
}
