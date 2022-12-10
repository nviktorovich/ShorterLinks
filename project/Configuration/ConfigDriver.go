package Configuration

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	Prefix     string
	DBInit     string
	DriverName string
)

// Config структура, предназначена для инициализации параметров (констант) из
// конфигурационного файла для исполнения программой
type Config struct {
	Prefix     string `yaml:"prefix"`
	DBInit     string `yaml:"db_init"`
	DriverName string `yaml:"driver_name"`
}

func NewConfig() *Config {
	return &Config{"", "", ""}
}

func ReadConfig() {
	cfg := NewConfig()
	file, err := os.ReadFile("project/Configuration/Config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	Prefix = cfg.Prefix
	DBInit = cfg.DBInit
	DriverName = cfg.DriverName
}
