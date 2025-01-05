package internals
import (
	"fmt"
	"os"
)

// Function checks to see if our full path to the exfil directory already exists, if it does the program will warn
// the user that there is a previous run from UnixCollector, we dont want to blow over previous collection
// the program will exit after warning the user 
//
// Return: bool -> if the exfil directory exists or not
func CheckExfilExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		fmt.Println("[!] Cannot determine if", path, "exists")
		fmt.Println("\tError Returned:", err)
		return true
	} else {
		// oh no something exists with our same name, lets see what it is
		info, _ := os.Stat(path)
		if info.IsDir() {
			fmt.Println("[!] Exfil directory", path, "exists, wont overwrite")
			return true
		} else {
			fmt.Println("[!] Exfil directory", path, "exists, however its a file")
			return true
		}
	}

}

// Function gets the cwd and then creates the exfil directory root where all exfiled files will be stored before 
// being compressed down. It will function as our / for the mini file system we will be building
//
// Return: string -> the path to the created exfiled directory 
func GetExfilDirectory() string {
	exfilDirectory := "exdata"
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("[!] Failed to determine cwd", err)
		os.Exit(1)
	}

	if CheckExfilExists(pwd+"/"+exfilDirectory) {
		// errors printed in above function, not needed here
		os.Exit(3)
	}

	err = os.Mkdir(pwd+"/"+exfilDirectory, 0644)
	if err != nil {
		fmt.Println("[!] Failed to create exfil directory")
		os.Exit(2)
	}
	// the full path to our exfil directory
	return pwd+"/"+exfilDirectory

}