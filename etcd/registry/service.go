package registry

import (
	"encoding/json"
	registry2 "github.com/jageros/hawos/registry"
)

func marshal(si *registry2.ServiceInstance) (string, error) {
	data, err := json.Marshal(si)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func unmarshal(data []byte) (si *registry2.ServiceInstance, err error) {
	err = json.Unmarshal(data, &si)
	return
}
