build:
	go build -ldflags="-s -w" sls.go
	$(if $(shell command -v upx || which upx), upx sls)

mac:
	GOOS=darwin go build -ldflags="-s -w" -o sls-darwin sls.go
	$(if $(shell command -v upx || which upx), upx sls-darwin)

win:
	GOOS=windows go build -ldflags="-s -w" -o sls.exe sls.go
	$(if $(shell command -v upx || which upx), upx sls.exe)

linux:
	GOOS=linux go build -ldflags="-s -w" -o sls-linux sls.go
	$(if $(shell command -v upx || which upx), upx sls-linux)