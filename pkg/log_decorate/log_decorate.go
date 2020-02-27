package log_decorate

import (
	`log`
)

type CallPhone func(number string) (err error)
type MailCompOrPhone func(number string, text string) (err error)

func LogCallPhone(fn CallPhone) (err CallPhone) {
	return func(number string) (err error) {
		log.Println("Start call phone")
		err = fn(number)
		if err != nil {
			log.Println("Error Call", err)
			return
		}
		log.Println("Success call for number", number)
		return
	}
}

func LogMailCompOrPhone(fn MailCompOrPhone) (err MailCompOrPhone) {
	return func(number string, text string) (err error) {
		log.Println("Start call phone")
		err = fn(number, text)
		if err != nil {
			log.Println("Error send message", err)
			return
		}
		log.Println("Success send message", number, text)
		return
	}
}
