# shellcode 免杀loader 生成

## 加密方式支持
- AES
- DES
- RC4
- CAST5
- IDEA

## 加载方式支持
- EARLY
- HEAP
- FIBER
- RTL
- UUID
- VIRT

## 沙箱检测支持
- 检测CPU
- 检测临时文件
- 检测虚拟机文件
- 检测启动时间
- 检测内存

## 生成方式
```
GOOS=windows
GOARCH=amd64
ENCODE := aes
CALL := early
KEY := yT6kL8kK6jJ3aO2e
SOURCE := cRyLMeJRSVDT3Pt80gFE6wfQsHS7m30uOzfbB5CH36g=
CURRENT_TIME_STAMP := $(shell date +%s)
make
```
