package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	file_reader("a.txt")
}

func file_reader(p string) {
	file, _ := os.Open(p)
	defer file.Close()
	for {
		<-time.After(time.Second * 2)
		data := make([]byte, 1<<16)
		switch _, err := file.Read(data); err {
		case nil:
			log.Println(strings.Split(string(data), "\n"))
			log.Println(len(strings.Split(string(data), "\n")))
		case io.EOF:
			log.Println(err)
		default:
			fmt.Println(err)
			return
		}
	}
}
