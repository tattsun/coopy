production:
	go-bindata -pkg="views" -o views/static.go static/
	go install github.com/tattsun/coopy

debug:
	go-bindata -pkg="views" -debug -o views/static.go static/
	go install github.com/tattsun/coopy

watch:
	go-bindata -pkg="views" -debug -o views/static.go static/
	COOPY_MYSQL_HOST=localhost:3306 COOPY_MYSQL_USER=coopyuser COOPY_MYSQL_PASSWORD=hogehoge COOPY_MYSQL_DATABASE=coopy fresh
