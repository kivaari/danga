package backend

import (
	"fmt"
	"os/exec"
)

// Launcher содержит функции для управления запуском игры
type Launcher struct{}

// StartGame запускает Minecraft с параметрами из config/settings.json
func (l *Launcher) StartGame() error {
	cmd := exec.Command("java", "-jar", "assets/minecraft.jar") // Укажите правильный путь к файлу
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("ошибка запуска Minecraft: %w", err)
	}
	return nil
}

// Add другие функции: Проверка обновлений, загрузка ресурсов, установка модов и т.д.
