package minecraft

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Callback struct {
	SetStatus   func(string)
	SetProgress func(int)
	SetMax      func(int)
}

type VersionData struct {
	ID         string `json:"id"`
	Assets     string `json:"assets"`
	AssetIndex struct {
		URL  string `json:"url"`
		Sha1 string `json:"sha1"`
	} `json:"assetIndex"`
	Libraries []Library `json:"libraries"`
}

type Library struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Action string `json:"action"`
}

func downloadFile(url string, dest string, _ *Callback) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to download file")
	}

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func installLibraries(libraries []Library, path string, callback *Callback) error {
	libraryPath := filepath.Join(path, "libraries")
	callback.SetStatus("Downloading Libraries")
	callback.SetMax(len(libraries))

	for i, lib := range libraries {
		libPath := filepath.Join(libraryPath, lib.Name)
		if _, err := os.Stat(libPath); os.IsNotExist(err) {
			if err := downloadFile(lib.URL, libPath, callback); err != nil {
				fmt.Printf("Failed to download library: %s\n", lib.Name)
			}
		}
		callback.SetProgress(i + 1)
	}
	return nil
}

func installAssets(data VersionData, path string, callback *Callback) error {
	assetsPath := filepath.Join(path, "assets", "indexes", data.Assets+".json")
	callback.SetStatus("Downloading Assets")

	err := downloadFile(data.AssetIndex.URL, assetsPath, callback)
	if err != nil {
		return err
	}

	file, err := os.Open(assetsPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var assetsData map[string]interface{}
	if err := json.NewDecoder(file).Decode(&assetsData); err != nil {
		return err
	}

	objects := assetsData["objects"].(map[string]interface{})
	callback.SetMax(len(objects))
	count := 0
	for _, value := range objects {
		object := value.(map[string]interface{})
		hash := object["hash"].(string)
		assetURL := fmt.Sprintf("https://resources.download.minecraft.net/%s/%s", hash[:2], hash)
		assetPath := filepath.Join(path, "assets", "objects", hash[:2], hash)
		if err := downloadFile(assetURL, assetPath, callback); err != nil {
			fmt.Printf("Failed to download asset: %s\n", hash)
		}
		count++
		callback.SetProgress(count)
	}
	return nil
}

func InstallMinecraftVersion(versionID string, minecraftPath string, callback *Callback) error {
	versionFile := filepath.Join(minecraftPath, "versions", versionID, versionID+".json")

	var versionData VersionData
	if _, err := os.Stat(versionFile); os.IsNotExist(err) {
		versionManifestURL := "https://launchermeta.mojang.com/mc/game/version_manifest_v2.json"
		resp, err := http.Get(versionManifestURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var manifest struct {
			Versions []struct {
				ID  string `json:"id"`
				URL string `json:"url"`
			} `json:"versions"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&manifest); err != nil {
			return err
		}

		for _, v := range manifest.Versions {
			if v.ID == versionID {
				if err := downloadFile(v.URL, versionFile, callback); err != nil {
					return err
				}
				break
			}
		}
	}

	file, err := os.Open(versionFile)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&versionData); err != nil {
		return err
	}

	if err := installLibraries(versionData.Libraries, minecraftPath, callback); err != nil {
		return err
	}
	if err := installAssets(versionData, minecraftPath, callback); err != nil {
		return err
	}

	callback.SetStatus("Installation Complete")
	return nil
}
