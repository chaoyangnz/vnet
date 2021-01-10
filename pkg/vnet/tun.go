package vnet

import (
	"fmt"
	"github.com/songgao/water"
	"github.com/vishvananda/netlink"
	"os/exec"
	"runtime"
)

type TunInterface struct {
	tun *water.Interface
	name string
	address string
}

func NewTunInterface(name string, address string) (*TunInterface, error) {
	iface := &TunInterface{
		name: name,
		address: address,
	}

	ifconfig := water.Config{
		DeviceType: water.TUN,
	}
	ifconfig.Name = name

	ifce, err := water.New(ifconfig)
	if err != nil {
		return nil, err
	}

	iface.tun = ifce
	return iface, nil
}

func (iface *TunInterface) Up() error {
	switch runtime.GOOS {
	case "linux", "darwin":
		out, err := execCmd("ifconfig", []string{iface.tun.Name(), "up"})
		if err != nil {
			return fmt.Errorf("ifconfig fail: %s %v", out, err)
		}
		if tun, err := netlink.LinkByName(iface.name); err == nil {
			addr, _ := netlink.ParseAddr(iface.address)
			netlink.AddrAdd(tun, addr)
			fmt.Printf("%s add address: %s \n", iface.name, iface.address)
		}

	default:
		return fmt.Errorf("unsupported: %s %s", runtime.GOOS, runtime.GOARCH)

	}

	return nil
}

func (iface *TunInterface) Read() ([]byte, error) {
	buf := make([]byte, 1024)
	n, err := iface.tun.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf[:n], nil
}

func (iface *TunInterface) Write(buf []byte) (int, error) {
	return iface.tun.Write(buf)
}

func (iface *TunInterface) Close() {
	iface.tun.Close()
}

func execCmd(cmd string, args []string) (string, error) {
	b, err := exec.Command(cmd, args...).CombinedOutput()
	return string(b), err
}
