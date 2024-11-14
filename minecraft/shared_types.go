package minecraft

type ClientJsonRule struct {
	Action   string            `json:"action"`
	Os       map[string]string `json:"os"`
	Features map[string]bool   `json:"features"`
}

type ClientJsonArgumentRule struct {
	CompatibilityRules []ClientJsonRule `json:"compatibilityRules,omitempty"`
	Rules              []ClientJsonRule `json:"rules,omitempty"`
	Value              interface{}      `json:"value"`
}

type ClientJsonAssetIndex struct {
	ID        string `json:"id"`
	Sha1      string `json:"sha1"`
	Size      int    `json:"size"`
	TotalSize int    `json:"totalSize"`
	URL       string `json:"url"`
}

type ClientJsonDownloads struct {
	Sha1 string `json:"sha1"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}
