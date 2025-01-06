package internals

import (
	"fmt"
	"io"
	"os"
	"strings"
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

	err = os.Mkdir(pwd+"/"+exfilDirectory, 0755)
	if err != nil {
		fmt.Println("[!] Failed to create exfil directory")
		os.Exit(2)
	}
	// the full path to our exfil directory
	return pwd+"/"+exfilDirectory

}

// Function takes in the preprocess list of found files and ensures we can access each file
// if we cannot access the file it is stripped from our file array
//
// Return: -> an array of accessible files
func FilterExistingFiles(paths []string) []string {
	var existingFiles []string
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			// we can stat the file but can we open it for a later copy operation
			_, err := os.OpenFile(path, os.O_RDONLY, 0644)
			if err == nil {
				existingFiles = append(existingFiles, path)
			} else {
				continue
			}	
		}
	}
	return existingFiles
}

func RebuildDirs(fileParts []string, exfilDirectory string) {
	filePartsNoFilename := fileParts[:len(fileParts)-1]
	var exfilDirectoryBuilder []string
	for _, i := range filePartsNoFilename {
		exfilDirectoryBuilder = append(exfilDirectoryBuilder, i)
		_, err := os.Stat(exfilDirectory + "/" + strings.Join(exfilDirectoryBuilder, "/"))
		if os.IsNotExist(err) {
			os.Mkdir(exfilDirectory + "/" + strings.Join(exfilDirectoryBuilder, "/"), 0755)
		} else {
			continue
		}
	}
}

func CopyFiles(toCollect []string, exfilDirectory string) {
	for _, file := range toCollect {
		fileParts := strings.Split(file, "/")
		
		RebuildDirs(fileParts, exfilDirectory)
		
		sourceFile, err := os.OpenFile(file, os.O_RDONLY, 0755)
		if err != nil {
			fmt.Println("[!] Error accessing:", file)
			fmt.Println("\tError Returned:", err)
			// pass onto the next file in the array 
			continue
		}
		destFile, err := os.Create(exfilDirectory+file)
		if err != nil {
			fmt.Println("[!] Error creating file", exfilDirectory+file)
			fmt.Println("\tError Returned:", err)
			// pass onto the next file in the array 
			continue
		}
		_, err = io.Copy(destFile, sourceFile)
		if err != nil {
			fmt.Println("[!] Error copying file", sourceFile, "->", destFile)
			fmt.Println("\tError Returned:", err)
			continue
		}
	}
}