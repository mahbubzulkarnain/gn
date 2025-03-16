package config

import "encoding/json"

type PathsConfig []string

func (c *PathsConfig) UnmarshalJSON(data []byte) error {
	if string(data[0]) == `[` {
		var out []string
		if err := json.Unmarshal(data, &out); err != nil {
			return nil
		}
		*c = out
		return nil
	}
	var out string
	if err := json.Unmarshal(data, &out); err != nil {
		return nil
	}
	*c = []string{out}
	return nil
}

func (c *PathsConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var out []string
	if sliceErr := unmarshal(&out); sliceErr != nil {
		var ele string
		if strErr := unmarshal(&ele); strErr != nil {
			return strErr
		}
		out = []string{ele}
	}

	*c = out
	return nil
}
