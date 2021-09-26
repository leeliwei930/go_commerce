package inputs

import (
	"encoding/json"
	"io"
)

// Decode JSON in for io.ReadCloser to struct object
func BindFromJSON(read io.ReadCloser, obj interface{}) error {
	return json.NewDecoder(read).Decode(obj)
}

func BindFromJSONX(read io.ReadCloser, obj interface{}) {
	if err := json.NewDecoder(read).Decode(obj); err != nil {
		panic(err)
	}
}
