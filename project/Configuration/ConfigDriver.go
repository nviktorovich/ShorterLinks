package Configuration

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	Prefix       string
	DBInit       string
	DriverName   string
	MainPage     string
	TitlePath    string
	ShortPage    string
	ShortPath    string
	RedirectPage string
	Address      string
)

// Config структура, предназначена для инициализации параметров (констант) из
// конфигурационного файла для исполнения программой
type Config struct {
	Prefix       string `yaml:"prefix"`
	DBInit       string `yaml:"db_init"`
	DriverName   string `yaml:"driver_name"`
	MainPage     string `yaml:"main_page"`
	TitlePath    string `yaml:"title_path"`
	ShortPage    string `yaml:"short_page"`
	ShortPath    string `yaml:"short_path"`
	RedirectPage string `yaml:"redirect_page"`
	Address      string `yaml:"address"`
}

func NewConfig() *Config {
	return &Config{
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}
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
	MainPage = cfg.MainPage
	TitlePath = cfg.TitlePath
	ShortPage = cfg.ShortPage
	ShortPath = cfg.ShortPath
	RedirectPage = cfg.RedirectPage
	Address = cfg.Address
}
