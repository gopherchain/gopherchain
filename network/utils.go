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

func fileEmpty(filename string) bool {
	file, err := os.Stat(filename)
	if err != nil {
		log.Panic(err)
	}

	if file.Size() == 0 {
		return true
	}

	return false
}

func defaultNodeConfig(filename string) {
	fmt.Println("Creating default nodes file")
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString("178.128.134.203:3000\n"); err != nil {
		log.Println(err)
	}
}
