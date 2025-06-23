package file

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"path"
	"sync"

	"gopkg.in/yaml.v3"
)

var groups map[string]int
var shares map[string]int
var singletonOnce sync.Once
var backend *Backend
var backendLock sync.RWMutex

func Initialize() (_ *Backend, err error) {
	singletonOnce.Do(func() {
		backend = &Backend{sharesPath: viper.GetString("shares_backend.file.file")}
		if err = backend.loadShares(); err != nil {
			return
		}
		return
	})

	return backend, err
}

func (b *Backend) loadShares() (err error) {
	backendLock.Lock()
	defer backendLock.Unlock()

	var yamlFile []byte

	if yamlFile, err = os.ReadFile("./data/shares.yml"); err != nil {
		return err
	}
	if err = yaml.Unmarshal(yamlFile, b); err != nil {
		return err
	}

	// cache to map
	shares = make(map[string]int)
	groups = make(map[string]int)

	for idx, g := range b.Groups {
		groups[g.ID] = idx
	}
	for idx, s := range b.Shares {
		shares[s.ShareID] = idx
	}

	slog.Info("backend shares loaded", slog.Int("shares", len(shares)), slog.Int("groups", len(groups)))
	return nil
}

func (b *Backend) GetFilePath(groupID string, shareID string) (string, error) {
	backendLock.RLock()
	defer backendLock.RUnlock()

	_share := b.getShareByID(shareID)
	_group := b.getGroupByID(groupID)

	if _group != nil && _share != nil {
		fullFilepath := path.Join(_group.RootPath, _share.Path)
		return fullFilepath, nil
	}

	return "", fmt.Errorf("share not found")
}

func (b *Backend) getShareByID(shareID string) *share {
	backendLock.RLock()
	defer backendLock.RUnlock()

	if sIdx, exists := shares[shareID]; exists {
		_share := b.Shares[sIdx]
		if _share.ShareID == shareID {
			return &_share
		}
	}
	return nil
}

func (b *Backend) getGroupByID(groupID string) *group {
	backendLock.RLock()
	defer backendLock.RUnlock()

	if gIdx, exists := groups[groupID]; exists {
		_group := b.Groups[gIdx]
		if _group.ID == groupID {
			return &_group
		}
	}
	return nil
}
