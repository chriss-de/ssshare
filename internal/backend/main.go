package backend

import (
	"sync"

	"github.com/chriss-de/ssshare/internal/backend/file"
)

var singletonOnce sync.Once
var backend Backend

func Initialize() (err error) {
	singletonOnce.Do(func() {
		backend, err = file.Initialize()
	})

	return err
}

func GetFilePath(groupID string, fileID string) (string, error) {
	return backend.GetFilePath(groupID, fileID)
}
