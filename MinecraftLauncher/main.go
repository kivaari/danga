package main

import (
	"embed"
	"encoding/json"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// Settings хранит настройки лаунчера
type Settings struct {
	JavaPath      string `json:"javaPath"`
	MinecraftPath string `json:"minecraftPath"`
	// Другие поля настроек
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "MinecraftLauncher",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// LoadSettings загружает настройки из config/settings.json
func LoadSettings() (*Settings, error) {
	file, err := os.Open("config/settings.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	settings := &Settings{}
	err = json.NewDecoder(file).Decode(settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}

// SaveSettings сохраняет настройки в config/settings.json
func SaveSettings(settings *Settings) error {
	file, err := os.Create("config/settings.json")
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(settings)
}
