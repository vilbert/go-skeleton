package config

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

var (
	config *Config
)

const (
	development = "development"
	staging     = "staging"
	production  = "production"
)

type option struct {
	configFile string
}

// Init ...
func Init(opts ...Option) error {
	opt := &option{
		configFile: getDefaultConfigFile(),
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}

	out, err := ioutil.ReadFile(opt.configFile)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(out, &config)
}

// Option ...
type Option func(*option)

// WithConfigFile ...
func WithConfigFile(file string) Option {
	return func(opt *option) {
		opt.configFile = file
	}
}

func getDefaultConfigFile() string {
	env := "development"
	configPath := filepath.Join(getRootDir(), "files/etc/skeleton")
	// For Kubernetes namespaces
	// namespace, _ := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	// env = string(namespace)
	//

	switch env {
	case development:
		// Docker Env
		if os.Getenv("GOPATH") == "" {
			configPath = "./skeleton.development.yaml"
		} else {
			configPath = filepath.Join(configPath, "/skeleton.development.yaml")
		}
	case staging:
		configPath = filepath.Join(configPath, "/skeleton.staging.yaml")
	case production:
		configPath = filepath.Join(configPath, "/skeleton.production.yaml")
	}
	return configPath
}

// Get ...
func Get() *Config {
	return config
}

func getRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Join(filepath.Dir(d), "..")
}
