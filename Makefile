production:
	go-bindata -pkg="views" -o views/static.go static/...
	go install github.com/tattsun/coopy

debug:
	go-bindata -pkg="views" -debug -o views/static.go static/...
	go install github.com/tattsun/coopy

install-dev:
	go get
	go get github.com/pilu/fresh
	go install github.com/pilu/fresh
	go get github.com/jteeuwen/go-bindata
	go install github.com/jteeuwen/go-bindata

watch:
	go-bindata -pkg="views" -debug -o views/static.go static/...
	COOPY_MYSQL_HOST=localhost:3306 COOPY_MYSQL_USER=coopyuser COOPY_MYSQL_PASSWORD=hogehoge COOPY_MYSQL_DATABASE=coopy fresh

test:
	go test ./...
