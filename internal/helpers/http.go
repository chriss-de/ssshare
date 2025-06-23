package helpers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v3"
)

type H map[string]any

func convertError(data any) any {
	if data != nil {
		if dataErr, ok := data.(error); ok {
			return H{"error": dataErr}
		}
	}
	return data
}

func WriteJSON(w http.ResponseWriter, status int, data any) (_ int, err error) {
	data = convertError(data)
	var jsonBytes []byte
	if jsonBytes, err = json.Marshal(data); err != nil {
		return 0, err
	}
	return WriteData(w, status, "application/json; charset=utf-8", jsonBytes)
}

func WriteYAML(w http.ResponseWriter, status int, data any) (_ int, err error) {
	data = convertError(data)
	var yamlBytes []byte
	if yamlBytes, err = yaml.Marshal(data); err != nil {
		return 0, err
	}
	return WriteData(w, status, "application/yaml; charset=utf-8", yamlBytes)
}

func WriteData(w http.ResponseWriter, status int, contentType string, data []byte) (_ int, err error) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)

	if data != nil && status != http.StatusNoContent {
		return w.Write(data)
	}
	return 0, nil
}
