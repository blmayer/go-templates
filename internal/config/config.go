package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	configFilePath = "/program/main.conf"
)

var (
	configPath string
)

// Database represents a database
type Database struct {
	ConnString string
	Schema     string
}

// Config is the main struct for configuration
type Config struct {
	ConnString string
	Database   string
}

func init() {
	// Loading
	var err error
	configPath, err = os.UserConfigDir()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	configPath += configFilePath
}

// Load reads the configuration file
func Load() (c Config, err error) {
	// Defaults
	c.ConnString = "mongodb+srv://localhost"
	c.Database = "program"

	configFile, err := os.Open(configPath)
	if err != nil {
		return
	}
	defer configFile.Close()
	err = json.NewDecoder(configFile).Decode(&c)
	if err != nil {
		return
	}

	return c, nil
}

// Configure runs the configuration prompt
func Configure() {
	configFile, err := os.OpenFile(configPath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
	defer configFile.Close()

	c := Config{}

	// Connection setup
	var ans string
	println("Configure database connection? [y/N]")
	fmt.Scanf("%s", &ans)
	if ans == "y" {
		println("Enter the connection string:")
		fmt.Scanf("%s", &c.ConnString)

		println("Enter the database name:")
		fmt.Scanf("%s", &c.Database)
		println("Done.")
	}

	// Write
	enc := json.NewEncoder(configFile)
	enc.SetIndent("", "\t")

	err = enc.Encode(c)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}

// Save writes the current config to disk
func Save(c Config) error {
	configFile, err := os.OpenFile(configPath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer configFile.Close()
	enc := json.NewEncoder(configFile)
	enc.SetIndent("", "\t")

	return enc.Encode(c)
}
