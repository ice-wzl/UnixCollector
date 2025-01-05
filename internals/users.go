package internals 
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function gets all the users home directories 
//
// Return: []string -> and array of all the users home directories
func getUsersHomedir() []string {
	// array to hold all the users on the system
	var users []string
	badUsers := []string{"/usr/sbin/nologin", "/bin/sync", "/bin/false"}

	if _, err := os.Stat("/etc/passwd"); err != nil {
		fmt.Printf("[!] Error getting users: %v\n", err)
		return users
	}
	file, err := os.OpenFile("/etc/passwd", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("[!] Error opening file: %v\n", err)
		return users
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ":")
		if len(fields) < 7 {
			// pass over lines that dont have 7 fields
			continue
		}
		isBadUsers := false
		for _, b := range badUsers {
			if fields[6] == b {
				isBadUsers = true
				break
			}
		}
		if !isBadUsers {
			users = append(users, fields[5])
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("[!] Error reading file: %v\n", err)
	}

	return users
}