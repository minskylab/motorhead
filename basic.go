package motorhead

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func ReadScript(scriptPath string) {
	data, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal()
}
