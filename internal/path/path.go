package path

import (
    "os"
    "path/filepath"
    "runtime"
)

func UserConfigDir() string {
    switch runtime.GOOS {
    case "windows":
        appDataPath := os.Getenv("APPDATA")
        if appDataPath == "" {
            h, _ := os.UserHomeDir()
            appDataPath = filepath.Join(h, "AppData", "Roaming")
        }
        return filepath.Join(appDataPath, "rhsmctl")
    default:
        xdgConfigPath := os.Getenv("XDG_CONFIG_HOME")
        if xdgConfigPath == "" {
            h, _ := os.UserHomeDir()
            xdgConfigPath = filepath.Join(h, ".config")
        }
        return filepath.Join(xdgConfigPath, "rhsmctl")
    }
}

func UserConfigFile() string {
	return filepath.Join(UserConfigDir(), "config.yml")
}
