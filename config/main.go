package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
)

// Settings contains the configuration details
type Settings struct {
	LogLevel                                       logrus.Level
	Debug, Simulate, StoredToken, Verbose, Version bool
	ConfigComment, EnvSSHDefaultUsername, EnvSSHIdentityFile,
	EnvSSHUsername, SecretPath, Show, StoragePath, User, VaultAddr, VaultToken string
}

var (
	once     sync.Once
	settings Settings

	/*
		The following support overrides during builds, which can be done
		by setting ldflags, e.g.

		`-ldflags "-X github.com/cezmunsta/ssh_ms/config.EnvSSHUserName=xxx"`

	*/

	// EnvBasePath is the parent location used to prefix storage paths
	EnvBasePath = filepath.Join(os.Getenv("HOME"), ".ssh", "cache")

	// EnvSSHDefaultUsername sets the default used in connections
	EnvSSHDefaultUsername = os.Getenv("USER")

	// EnvSSHUsername is used to authenticate with SSH
	EnvSSHUsername = "SSH_MS_USERNAME"

	// EnvSSHIdentityFile is used for SSH authentication
	EnvSSHIdentityFile = filepath.Join("~", ".ssh", "id_ed25519")

	// SecretPath is the location used for connection manangement
	SecretPath = "secret/ssh_ms"
)

// ToJSON returns the config in JSON format
func (s *Settings) ToJSON() string {
	data, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(data)
}

// GetConfig returns an instance of Settings
// ensuring that only one instance is ever returned
func GetConfig() *Settings {
	once.Do(func() {
		settings = Settings{
			ConfigComment:         "",
			EnvSSHDefaultUsername: EnvSSHDefaultUsername,
			EnvSSHIdentityFile:    EnvSSHIdentityFile,
			EnvSSHUsername:        EnvSSHUsername,
			LogLevel:              logrus.WarnLevel,
			SecretPath:            SecretPath,
			Simulate:              false,
			StoragePath:           EnvBasePath,
			StoredToken:           false,
		}
	})
	return &settings
}
