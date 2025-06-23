package backend

type Backend interface {
	GetFilePath(groupID string, shareID string) (string, error)
}
