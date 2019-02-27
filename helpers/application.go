package helpers

import "log"

func HandleError(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}
