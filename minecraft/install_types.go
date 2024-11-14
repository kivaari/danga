package minecraft

// AssetsJsonObject represents the hash and size of a specific asset.
type AssetsJsonObject struct {
	Hash string `json:"hash"`
	Size int    `json:"size"`
}

// AssetsJson represents the JSON structure for Minecraft assets.
type AssetsJson struct {
	Objects map[string]AssetsJsonObject `json:"objects"`
}
