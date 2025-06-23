package file

type Backend struct {
	Groups []group `json:"groups"`
	Shares []share `json:"shares"`
}

type group struct {
	ID       string `yaml:"id"`
	Alias    string `yaml:"alias"`
	RootPath string `yaml:"root_path"`
}

type share struct {
	GroupID string `yaml:"group_id"`
	ShareID string `yaml:"share_id"`
	Path    string `yaml:"path"`
}
