package transform

import (
	"encoding/json"
	"testing"

	simplejson "github.com/bitly/go-simplejson"
)

const testJSONInput = `{"rating":{"example":{"value":3},"primary":{"value":3}}}`

func getConfig(spec string, require bool) Config {
	var f map[string]interface{}
	json.Unmarshal([]byte(spec), &f)
	return Config{Spec: &f, Require: require}
}

func getTransformTestWrapper(tform func(spec *Config, data *simplejson.Json) error, cfg Config, input string) (string, error) {
	json, e := simplejson.NewJson([]byte(input))
	if e != nil {
		return "", e
	}
	e = tform(&cfg, json)
	if e != nil {
		return "", e
	}
	tmp, e := json.MarshalJSON()
	if e != nil {
		return "", e
	}
	return string(tmp), nil
}

func getTransformTestWrapperRaw(tform func(spec *Config, data []byte) ([]byte, error), cfg Config, input string) (string, error) {
	output, e := tform(&cfg, []byte(input))
	if e != nil {
		return "", e
	}
	return string(output), nil
}

func TestBookend(t *testing.T) {
	input := []byte(`"foo", "bar"`)
	expected := []byte(`["foo", "bar"]`)

	result := bookend(input, '[', ']')
	if string(result) != string(expected) {
		t.Error("Bookend result does not match expectation.")
		t.Log("Expected: ", expected)
		t.Log("Actual:   ", result)
		t.FailNow()
	}

	input = []byte("fooString")
	expected = []byte(`"fooString"`)
	result = bookend(input, '"', '"')
	if string(result) != string(expected) {
		t.Error("Bookend result does not match expectation.")
		t.Log("Expected: ", expected)
		t.Log("Actual:   ", result)
		t.FailNow()
	}
}
