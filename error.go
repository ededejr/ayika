package ayika

import "log"

func throwOnError(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}
