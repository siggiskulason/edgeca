package config

import (
	"os"
)

func GetDefaultTLSHost() string {
	hostName, _ := os.Hostname()
	return hostName
}
func GetDefaultTLSPort() int {
	return 50025
}
func GetDefaultTLSCertDir() string {
	homeDir, _ := os.UserHomeDir()
	configDir := homeDir + "/.edgeca"
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		_ = os.Mkdir(configDir, 0755)
	} else {
	}

	defaultTLSCertDir := configDir + "/certs"

	if _, err := os.Stat(defaultTLSCertDir); os.IsNotExist(err) {
		_ = os.Mkdir(defaultTLSCertDir, 0755)
	} else {
	}

	return defaultTLSCertDir
}
