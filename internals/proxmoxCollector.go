package internals

func ProxmoxConfigCollector() []string {
	paths := []string{
		"/etc/pve",
	}
	return paths
}
