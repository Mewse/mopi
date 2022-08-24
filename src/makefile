build:
	go mod download
	go build -o ../bin/mopi cmd/main.go
	$(eval version=$(shell sh -c 'cat version.txt 2>&1'))
	tar -zcf ../bin/mopi-${version}.tar.gz ../bin/mopi ./README.md ./CHANGES.md

build-docker:
	$(eval version=$(shell sh -c 'cat version.txt 2>&1'))
	docker build -t mopi:${version} .
	
test:
	go test

install:
	sudo cp ../bin/mopi /usr/local/bin/mopi

upload-dev:
	$(eval version=$(shell sh -c 'cat version.txt 2>&1'))
	docker tag mopi:${version} packages.mediakind.com/mpd-docker-dev/mediakind/mopi:${version}
	docker push packages.mediakind.com/mpd-docker-dev/mediakind/mopi:${version}

upload-prod:
	$(eval version=$(shell sh -c 'cat version.txt 2>&1'))
	docker tag mopi:${version} packages.mediakind.com/mpd-docker-prod/mediakind/mopi:${version}
	docker push packages.mediakind.com/mpd-docker-prod/mediakind/mopi:${version}