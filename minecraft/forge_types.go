package minecraft

type ForgeInstallProcessor struct {
	Sides     []string `json:"sides,omitempty"`
	Jar       string   `json:"jar,omitempty"`
	Classpath []string `json:"classpath,omitempty"`
	Args      []string `json:"args,omitempty"`
}

type ForgeInstallProfileInstall struct {
	ProfileName string `json:"profileName,omitempty"`
	Target      string `json:"target,omitempty"`
	Path        string `json:"path,omitempty"`
	Version     string `json:"version,omitempty"`
	FilePath    string `json:"filePath,omitempty"`
	Welcome     string `json:"welcome,omitempty"`
	Minecraft   string `json:"minecraft,omitempty"`
	MirrorList  string `json:"mirrorList,omitempty"`
	Logo        string `json:"logo,omitempty"`
}

type ForgeInstallProfile struct {
	Spec          int                          `json:"spec,omitempty"`
	Profile       string                       `json:"profile,omitempty"`
	Version       string                       `json:"version,omitempty"`
	Minecraft     string                       `json:"minecraft,omitempty"`
	ServerJarPath string                       `json:"serverJarPath,omitempty"`
	Data          map[string]map[string]string `json:"data,omitempty"`
	Processors    []ForgeInstallProcessor      `json:"processors,omitempty"`
	Libraries     []ClientJsonLibrary          `json:"libraries,omitempty"`
	Icon          string                       `json:"icon,omitempty"`
	Logo          string                       `json:"logo,omitempty"`
	MirrorList    string                       `json:"mirrorList,omitempty"`
	Welcome       string                       `json:"welcome,omitempty"`
	Install       ForgeInstallProfileInstall   `json:"install,omitempty"`
	VersionInfo   ClientJson                   `json:"versionInfo,omitempty"`
}
