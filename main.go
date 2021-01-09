package main

import (
	"fmt"
	"github.com/chaoyangnz/vnet"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	iface, err := vnet.NewTunInterface()
	if err != nil {
		fmt.Println("[E] new interface fail: ", err)
		return
	}

	//defer iface.Close()
	iface.Up()
	fmt.Printf("utun10 is up\n")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		iface.Close()
	}()

	for {
		buf, err := iface.Read()
		if err != nil {
			fmt.Printf("[E] read iface fail: %v\n", err)
			break
		}

		//fmt.Println(buf)
		version := buf[0] >> 4
		if version == 4 {
			datagram := vnet.NewIPV4Datagram(buf)
			datagram.Print()
		}
	}
}

