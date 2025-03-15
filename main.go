package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/shirou/gopsutil/host"
)

func getOS() string {
	return runtime.GOOS
}

func getHostName() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return hostname, nil
}

func getUser() string {
	return os.Getenv("USER")
}

func getKernelVersion() (string, error) {
	cmd := exec.Command("uname", "-r")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func formatUptime(seconds uint64) string {
	duration := time.Duration(seconds) * time.Second
	days := duration / (24 * time.Hour)
	duration -= days * 24 * time.Hour
	hours := duration / time.Hour
	duration -= hours * time.Hour
	minutes := duration / time.Minute
	return fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
}

func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, err
}

func main() {
	osName := getOS()
	hostname, err := getHostName()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}
	user := getUser()
	kernelVersion, err := getKernelVersion()
	if err != nil {
		fmt.Println("Error getting kernel version:", err)
		return
	}
	uptimeSeconds, err := host.Uptime()
	if err != nil {
		fmt.Println("Error getting uptime:", err)
		return
	}
	uptime := formatUptime(uptimeSeconds)
	ipAddress, err := GetOutboundIP()
	if err != nil {
		fmt.Println("Error getting IP address:", err)
		return
	}

	fmt.Printf("\033[38;5;9m\033[48;5;16m")
	myFigure := figure.NewFigure(hostname, "alligator", true)
	myFigure.Print()
	fmt.Printf("\033[0m")

	fmt.Printf("\nOS: %s\n", strings.Title(osName))
	fmt.Printf("Host: %s\n", hostname)
	fmt.Printf("User:  %s\n", user)
	fmt.Printf("Kernel Version: %s\n", kernelVersion)
	fmt.Printf("Uptime: %s\n", uptime)
	fmt.Printf("IP Address: %s\n", ipAddress)
}
