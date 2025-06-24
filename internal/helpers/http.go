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
			return H{"error": dataErr.Error()}
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

// GetScheme determinate http scheme from connection
func GetScheme(r *http.Request) string {
	if r.Header.Get("x-forwarded-proto") != "" {
		return r.Header.Get("x-forwarded-proto")
	}
	if r.URL.Scheme != "" {
		return r.URL.Scheme
	}
	if r.TLS != nil {
		return "https"
	}
	return "http" // "best" guess
}

// GetHost determinate http host from connection
func GetHost(r *http.Request) string {
	var (
		host = "127.0.0.1"
	)
	switch {
	case r.Header.Get("host") != "":
		host = r.Header.Get("host")
	case r.Header.Get("x-forwarded-host") != "":
		host = r.Header.Get("x-forwarded-host")
	case r.URL.Host != "":
		host = r.URL.Host
	case r.Host != "":
		host = r.Host
	default:
		host = "127.0.0.1"
	}
	return host
}
