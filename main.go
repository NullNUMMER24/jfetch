package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

func getUptime() (string, error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("bash", "-c", "uptime | sed 's/.*up \\([^,]*\\), .*/\\1/'")
	case "windows":
		cmd = exec.Command("net", "stats", "srv") // This is a workaround for Windows
	default:
		return "Uptime not available", nil
	}
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func getAsciiArt(osName string) string {
	switch osName {
	case "linux":
		return `Linux------------------------------------`
	case "darwin":
		return `Darwin-----------------------------------`
	case "windows":
		return `Windows----------------------------------`
	default:
		return "Unknown OS"
	}
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
	uptime, err := getUptime()
	if err != nil {
		fmt.Println("Error getting uptime:", err)
		return
	}

	fmt.Println(getAsciiArt(osName))
	fmt.Printf("OS: %s\n", strings.Title(osName))
	fmt.Printf("Host: %s\n", hostname)
	fmt.Printf("User:  %s\n", user)
	fmt.Printf("Kernel Version: %s\n", kernelVersion)
	fmt.Printf("Uptime: %s\n", uptime)
}
