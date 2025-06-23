package backend

import (
	"errors"
	"sync"

	"github.com/chriss-de/ssshare/internal/backend/file"

	"github.com/spf13/viper"
)

var singletonOnce sync.Once
var backend Backend

func Initialize() (err error) {
	singletonOnce.Do(func() {
		switch viper.GetString("shares.backend") {
		case "file":
			backend, err = file.Initialize()
		default:
			err = errors.New("backend not supported")
		}

	})

	return err
}

func GetFilePath(groupID string, fileID string) (string, error) {
	return backend.GetFilePath(groupID, fileID)
}
