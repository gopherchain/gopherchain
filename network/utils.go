package network

import (
	"fmt"
	"log"
	"os"
)

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func defaultNodeConfig(filename string) {
	fmt.Println("Creating default nodes file")
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString("204.48.19.128:3000\n"); err != nil {
		log.Println(err)
	}
}
