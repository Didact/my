package my

import (
	"log"
)

func Check(errs ...error) {
	for _, err := range errs {
		if err != nil {
			log.Fatal(err)
		}
	}
}
