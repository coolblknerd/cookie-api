package helper

import (
	"fmt"
	"reflect"
	"testing"

	c "github.com/coolblknerd/cookie-api/src/config"
)

var configs = SetUpConfigs()

func TestSetUpConfigs(t *testing.T) {
	if reflect.TypeOf(configs) != reflect.TypeOf(c.Configurations{}) {
		t.Error("Couldn't find any configs to load.")
	}
}

func TestDBURI(t *testing.T) {
	expected := fmt.Sprintf("mongodb://%s:%s@%s/%s", configs.Database.User, configs.Database.Password, configs.Database.Host, configs.Database.Name)
	uri := dbURI(configs)
	if uri != expected {
		t.Errorf("Expected: %s\nReceived: %s", expected, uri)
	}
}
