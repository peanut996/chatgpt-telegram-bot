package cfg

import (
	"flag"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	*EngineConfig `yaml:"engine"`
	*BotConfig    `yaml:"bot"`
}

type EngineConfig struct {
	EngineType string `yaml:"type"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
}

type BotConfig struct {
	BotType          string `yaml:"type"`
	TelegramBotToken string `yaml:"token"`
}

type RateLimiterConfig struct {
	Join        bool  `yaml:"join"`
	RemainCount bool  `yaml:"count"`
	Invite      bool  `yaml:"invite"`
	Capacity    int64 `yaml:"capacity"`
	Duration    int64 `yaml:"duration"`
}

func NewConfig() *Config {
	return &Config{}
}

func InitConfig() (*Config, error) {
	c := NewConfig()

	path := os.Getenv("CONFIG_PATH")

	if path == "" {
		flag.StringVar(&path, "c", "./config.yaml", "Your config file path")
		flag.Parse()
	}
	err := c.loadYaml(path)
	if err != nil {
		log.Fatalf("load config err: %v", err)
	}
	return c, nil
}

// Load config from config.yaml
func (c *Config) loadYaml(path string) error {
	yamlFile := path
	data, err := os.ReadFile(yamlFile)
	if nil != err {
		log.Printf("load local yaml err:%v path: %v\n", err, yamlFile)
		return err
	}

	err = yaml.Unmarshal([]byte(data), c)
	if nil != err {
		log.Printf("unmarshal local yaml err:%v path: %v\n", err, yamlFile)
		return err
	}
	return nil
}
