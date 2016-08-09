build:
	go build -o jkey .

build-osx:
	 GOOS=darwin GOARCH=amd64 go build -o jkey_darwin .

build-linux:
	GOOS=linux GOARCH=amd64 go build -o jkey_amd64 .
