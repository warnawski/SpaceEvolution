package configurate

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	DefaultConfigFilePath    = ""
	DefaultServerName        = ""
	DefaultServerDescription = ""
	DefaultServerPort        = 1414
	DefaultPublicIMG         = ""
)

type Conf struct {
	Path string `json:"path"`
	Data Data   `json:"data"`
}

type Data struct {
	ServerName        string `yaml:"server_name"`
	ServerDescription string `yaml:"server_description"`

	Port int `yaml:"port"`

	PublicIMG string `yaml:"public_img"`
}

func NewConf(Path string) *Conf {

	if Path == "" {
		log.Printf("Configuration file at path: %s, not found, starting installation of default config file", Path)
		Path = DefaultConfigFilePath
	}

	return &Conf{
		Path: Path,
	}
}

func (c *Conf) LoadConfig() error {
	data, err := os.ReadFile(c.Path)
	if err != nil {
		log.Printf("error reading configuration file: %s", err)
		return err
	}
	var dta Data

	err = yaml.Unmarshal(data, &dta)
	if err != nil {
		log.Printf("error parsing configuration file: %s", err)
		return err
	}

	c.configureValidate(&dta)
	c.Data = dta
	log.Printf("successful loading configuration file")

	return nil
}

func (c *Conf) configureValidate(data *Data) {

	if data.ServerName == "" {
		log.Printf("server name not found, default value set")
		data.ServerName = DefaultServerName
	}

	if data.ServerDescription == "" {
		log.Printf("the server description is empty, the default value is set")
		data.ServerDescription = DefaultServerDescription
	}

	if data.PublicIMG == "" {
		log.Printf("the server's custom icon was not found, starting installation of default icon")
		data.PublicIMG = DefaultPublicIMG
	}

	if data.Port < 1024 || data.Port > 65535 {
		log.Printf("the server's custom port was not found, starting installation of default server port: %d", data.Port)
		data.Port = DefaultServerPort
	}

}
