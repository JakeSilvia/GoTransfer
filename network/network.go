package network

import (
	"github.com/dean2021/go-nmap"
	"fmt"
)

type Device struct {
	DeviceName	string
	Ip 	string
}

func GetDevices() []Device {
	n := nmap.New()
	args := []string{"-sV", "-n", "-O", "-T4", "--open"}
	n.SetArgs(args...)
	n.SetPorts("0-65535")
	n.SetHosts("192.168.1.1/24")

	n.SetExclude("127.0.0.1")

	err := n.Run()
	if err != nil {
		fmt.Println("scanner failed:", err)
		return nil
	}

	result, err := n.Parse()
	if err != nil {
		fmt.Println("Parse scanner result:", err)
		return nil
	}
	fmt.Printf("Nmap: %+v", result)
	return nil
}
