package skillissue

import (
	"encoding/json"
	"os"
	"path"
)

type Config struct {
	configPath     string           `json:"-"`
	IgnoredPlayers ConfigPlayerList `json:"ignoredPlayers"`
	CommandPrefix  string           `json:"commandPrefix"`
	CoolPeople     ConfigPlayerList `json:"coolPeople"`
}

func LoadConfig(p string) (Config, error) {
	var c Config
	p = path.Clean(p)
	f, err := os.Open(p)
	if err != nil {
		return c, err
	}
	defer f.Close()
	c.IgnoredPlayers = make(ConfigPlayerList)
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
