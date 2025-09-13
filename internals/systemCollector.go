package internals

// /etc/httpd/conf/*.conf
// /etc/httpd/conf.d/*.conf
func HttpServerCollector() []string {
	paths := []string{
		"/etc/ssl",                   // SSL certificates
		"/etc/apache2/apache2.conf",  // Apache2 config
		"/etc/apache2/sites-enabled", // Apache2 sites-enabled
		"/etc/httpd/conf",            // httpd configs conf/ and conf.d/
		"/etc/httpd/conf.d",
		"/etc/nginx/conf.d", // nginx configs
		"/etc/phpmyadmin/config-db.php",
	}
	return paths
}

func SysinfoCollector() []string {
	paths := []string{
		"/etc/passwd", // User information
		"/etc/shadow",
		"/etc/group",    // Group information
		"/etc/hostname", // System hostname
		"/etc/hosts",    // Hosts file
		"/etc/fstab",    // Static file system information
	}
	return paths
}

func LogCollector() []string {
	paths := []string{
		"/var/log/auth.log", // Authentication logs (Linux-specific)
		"/var/log/secure",   // Secure logs (Red Hat/CentOS-specific)
		"/var/log/wtmp",     // Authentication logs, source ip
	}
	return paths
}

func SshSystemCollector() []string {
	paths := []string{
		"/etc/ssh",
		"/tmp/ssh-*", // Temporary SSH files
	}
	return paths
}

func KerberosCollector() []string {
	paths := []string{
		"/etc/krb5.conf",
		"/var/lib/krb5kdc",
	}
	return paths
}

func MiscSystemCollector() []string {
	paths := []string{
		"/etc/hosts.equiv", // rhosts system file
	}
	return paths
}
