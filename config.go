package agent

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

const defaultConfigName = "epg-agent.toml"

type Config struct {
	Sources []Source `toml:"source"`
	Dests   []Dest   `toml:"dest"`
}

type Source struct {
	Name string
	Path string
}

type Dest struct {
	Name     string
	Type     string
	Database DatabaseConfig `toml:"database"`
}

type DatabaseConfig struct {
	Host string
	Port int
}

func DecodeConfig(r io.Reader) (*Config, error) {
	var conf Config
	_, err := toml.DecodeReader(r, &conf)
	return &conf, err
}

func ReadConfigPath(filePath string) (*Config, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading file '%s': %s", filePath, err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("Error reading stat '%s': %s", filePath, err)
	}

	if fi.IsDir() {
		return ReadConfigPath(path.Join(filePath, defaultConfigName))
	}

	conf, err := DecodeConfig(f)
	if err != nil {
		return nil, fmt.Errorf("Error decoding file '%s': %s", filePath, err)
	}

	return conf, nil
}
