package skillissue

import (
	"encoding/json"
)

type ConfigPlayerList map[string]bool

func (i ConfigPlayerList) IsIgnored(name string) bool {
	return i[name]
}

func (i ConfigPlayerList) UnmarshalJSON(b []byte) error {
	var names []string
	if err := json.Unmarshal(b, &names); err != nil {
		return err
	}
	for _, name := range names {
		i[name] = true
	}
	return nil
}

func (i ConfigPlayerList) MarshalJSON() ([]byte, error) {
	arr := make([]string, 0, len(i))
	for k := range i {
		arr = append(arr, k)
	}
	return json.Marshal(arr)
}
