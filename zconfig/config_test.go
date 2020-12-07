package zconfig

import "testing"

func TestInit(t *testing.T) {
	config := make(map[string]interface{})
	config["test"] = true
	AddConfig(config)
}
