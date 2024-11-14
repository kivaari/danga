package minecraft

type RuntimeListJsonEntryManifest struct {
	Sha1 string `json:"sha1"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}

type RuntimeListJsonEntry struct {
	Availability map[string]int               `json:"availability"`
	Manifest     RuntimeListJsonEntryManifest `json:"manifest"`
	Version      map[string]string            `json:"version"`
}

type RuntimeListJson map[string]map[string][]RuntimeListJsonEntry

type PlatformManifestJsonFileDownloads struct {
	Sha1 string `json:"sha1"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}

type PlatformManifestJsonFile struct {
	Downloads  map[string]PlatformManifestJsonFileDownloads `json:"downloads,omitempty"`
	Type       string                                       `json:"type"`
	Executable bool                                         `json:"executable,omitempty"`
	Target     string                                       `json:"target,omitempty"`
}

type PlatformManifestJson struct {
	Files map[string]PlatformManifestJsonFile `json:"files"`
}
