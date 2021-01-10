package main

import (
	"fmt"
	"github.com/chaoyangnz/pkg/vnet"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	tif, err := vnet.NewTunInterface("tun0", "10.0.0.1/24")
	if err != nil {
		fmt.Println("[E] new interface 1 fail: ", err)
		return
	}
	tif.Up()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		tif.Close()
	}()

	for {
		buf, err := tif.Read()
		if err != nil {
			fmt.Printf("[E] read iface fail: %v\n", err)
			break
		}

		//fmt.Println(buf)
		version := buf[0] >> 4
		if version == 4 {
			onIPV4(vnet.NewIPv4Datagram(buf), tif)
		} else {
			fmt.Printf("Unsupported ip version: %d \n", version)
		}
	}
}

func onIPV4(datagram *vnet.IPv4Datagram, tif *vnet.TunInterface)  {
	datagram.Print()
}

