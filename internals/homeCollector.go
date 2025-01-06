package internals
import (
	"path/filepath"
)

func GitCollector(userHome string) []string {
	paths := []string{
		filepath.Join(userHome, ".git-credentials"),   // Git credentials
		filepath.Join(userHome, ".gitconfig"),          // Git global config
	}
	return paths
}

func SshCollector(userHome string) []string {
	paths := []string{
		filepath.Join(userHome, ".ssh"),                   // SSH keys
	}
	return paths
}

func CloudCollector(userHome string) []string {
	paths := []string{
		filepath.Join(userHome, ".aws"),                   // AWS credentials
		filepath.Join(userHome, ".config/gcloud"),         // Google Cloud config
		filepath.Join(userHome, ".azure"),                 // Azure config
	}
	return paths
}

func SqlCollector(userHome string) []string {
	paths := []string{
		filepath.Join(userHome, ".mysql_history"),         // MySQL history
		filepath.Join(userHome, ".psql_history"),          // Postgres history
		filepath.Join(userHome, ".dbshell"),               // DBShell history
		filepath.Join(userHome, ".pgpass"),                // Postgres Pass
		filepath.Join(userHome, ".config/sqlitebrowser/sqlitebrowser.conf"), 
	}
	return paths
}

func ShellCollector(userHome string) []string {
	paths := []string{
		filepath.Join(userHome, ".bash_history"),
		filepath.Join(userHome, ".zsh_history"),
		filepath.Join(userHome, ".ash_history"),
		filepath.Join(userHome, ".ksh_history"),
		filepath.Join(userHome, ".tcsh_history"),
	}
	return paths
}

func RcloneCollector(userHome string) []string {
	paths := []string{
		filepath.Join(userHome, ".config/rclone"),         // rclone backup configs
		filepath.Join(userHome, ".config/rclone_browser"), // rclone_browser configs
	}
	return paths
}

func RdpCollector(userHome string) []string {
	paths := []string{
		filepath.Join(userHome, ".config/freerdp"),        // Freerdp files
		filepath.Join(userHome, ".local/share/remmina"),   // Remmina RDP files
		filepath.Join(userHome, ".config/remmina"),        // Remmina RDP files
		filepath.Join(userHome, ".remmina"),               // Remmina RDP files
		filepath.Join(userHome, ".vnc"),                   // VNC files
	}
	return paths
}

	// /home/*/.docker/*/Dockerfile
	// /home/*/.docker/*/Containerfile
	// /home/*/.docker/*/*dockerenv*
	// /home/*/.docker/*/docker-compose.yml
func ContainerCollector(userHome string) []string {
	paths := []string{
		filepath.Join(userHome, ".docker"),                // Docker config
		filepath.Join(userHome, ".kube"),                  // Kubernetes config
	}
	return paths
}

func VpnCollector(userHome string) []string {
	paths := []string{
		filepath.Join(userHome, ".openvpn"),               // OpenVPN config
	}
	return paths
}

func EditorCollector(userHome string) []string {
	paths := []string{
		filepath.Join(userHome, ".python_history"),
		filepath.Join(userHome, ".wget-hsts"),
		filepath.Join(userHome, ".lesshst"),
		filepath.Join(userHome, ".viminfo"),
		filepath.Join(userHome, ".vimrc"),
		filepath.Join(userHome, ".profile"),               // User profile
	}
	return paths
}

func KeyringCollector(userHome string) []string {
	paths := []string{
		filepath.Join(userHome, ".gnupg"),                 // GPG keys
		filepath.Join(userHome, ".local/share/keyrings"),  // Keyrings
		filepath.Join(userHome, ".pki"), 				   // PKI files
	}
	return paths
}

func MiscCollector(userHome string) []string {
	paths := []string{
		filepath.Join(userHome, ".npmrc"),                 // NPM credentials
		filepath.Join(userHome, ".pypirc"),                // Python package repository credentials
		filepath.Join(userHome, ".netrc"),                 // Netrc (generic credentials)
		filepath.Join(userHome, ".config/teamviewer/clients.conf"), // teamviewer conf
		filepath.Join(userHome, ".filezilla"),
	}
	return paths
}

func MailCollector(userHome string) []string {
	paths := []string{
		filepath.Join(userHome, ".thunderbird"),
	}
	return paths
}
