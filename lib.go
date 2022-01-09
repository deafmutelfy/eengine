package main

import (
	"log"
	"os"
)

func check(err error, die bool) {
	if err != nil {
		log.Println(err)

		if die {
			log.Println("This error is critical, can't operate anymore.")
			os.Exit(1)
		}
	}
}