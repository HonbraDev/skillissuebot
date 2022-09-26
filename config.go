package skillissue

import (
	"encoding/json"
	"os"
	"path"
)

type Config struct {
	configPath    string           `json:"-"`
	CommandPrefix string           `json:"commandPrefix"`
	UncoolPeople  ConfigPlayerList `json:"uncoolPeople"`
	CoolPeople    ConfigPlayerList `json:"coolPeople"`
	BedCoords     Position         `json:"bedCoords"`
}

func LoadConfig(p string) (Config, error) {
	var c Config
	p = path.Clean(p)
	f, err := os.Open(p)
	if err != nil {
		return c, err
	}
	defer f.Close()
	c.UncoolPeople = make(ConfigPlayerList)
	c.CoolPeople = make(ConfigPlayerList)
	c.configPath = p
	err = json.NewDecoder(f).Decode(&c)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (c Config) Save(indent bool) error {
	f, err := os.Create(c.configPath)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	if indent {
		enc.SetIndent("", "  ")
	}
	return enc.Encode(c)
}
