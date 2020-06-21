binary: setup update generate build docker-build


build: setup
	GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -mod=vendor -ldflags "-X main.Version=`cat .version`" ./main.go;mv ./main ./dist/linux/bin/sc

setup:
	if [ -e ./dist ]; then rm -rf ./dist; fi; mkdir ./dist; mkdir -p ./dist/darwin/bin; mkdir -p ./dist/linux/bin


