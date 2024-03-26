package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
  DefaultConfigPath = ".config/fuzzy-cloner/config.yaml"
  DefaultVersion = "v1"
  DefaultCloneDir = "src"

  DefaultSource = "github"
  DefaultAuthMethod = "SSH"
  DefaultCredentialsPath = ".ssh/id_rsa.pub"
)

// TODO(dpe): handle the home filepath more gracefully
var homedir, _ = os.UserHomeDir()

type Config struct {
  Version string `yaml:"version"`
  CloneDir string  `yaml:"cloneDir"`
  GitOptions []GitOptions `yaml:"gitOptions"`
}

type GitOptions struct {
  Source string `yaml:"source"`
  AuthMethod string `yaml:"authMethod"`
  CredentialsPath string `yaml:"credentialPath"`
  CustomOptions CustomOptions `yaml:"customOptions"`
}

type CustomOptions map[string]string 

// New generates a new fuzzy-cloner config from a config file
func New() (*Config, error) {
  var config Config

  configPath := os.Getenv("FUZZY_CLONER_CONFIG_PATH")
  if configPath == "" {
   configPath = filepath.Join(homedir, DefaultConfigPath)
  }
  
  if _, err := os.Stat(configPath); os.IsNotExist(err) {
    fmt.Println("No config file set, using default parameters")
    return &config, nil
  }
  
  output, err := os.ReadFile(configPath)
  if err != nil {
    return &config, err
  }
  
  err = yaml.Unmarshal(output, &config)
  if err != nil {
    return &config, err
  }

  return &config, nil
}

func (c *Config) GetVersion() string {
  if c.Version == "" {
    return DefaultVersion
  }
  return c.Version
}
  
func (c *Config) GetCloneDir() string {
  if c.CloneDir == "" {
    return filepath.Join(homedir, DefaultCloneDir)
  }
  return c.CloneDir
}

// GetGitOptions outputs authentication options and custom options for all configured git sources
// and will default to github with ssh authentication if no values are set 
func (c *Config) GetGitOptions() []GitOptions {
  if c.GitOptions == nil  {
    return []GitOptions{{
        Source: DefaultSource,
        AuthMethod: DefaultAuthMethod,
        CredentialsPath: filepath.Join(homedir, DefaultCredentialsPath),
        CustomOptions: nil,
      },
    }
  }

  return c.GitOptions
}

