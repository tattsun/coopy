# coopy
[![wercker status](https://app.wercker.com/status/8440b7b9ecbbfc6575dcfbefc593641a/s "wercker status")](https://app.wercker.com/project/bykey/8440b7b9ecbbfc6575dcfbefc593641a)
[![Coverage Status](https://coveralls.io/repos/tattsun/coopy/badge.svg?branch=master)](https://coveralls.io/r/tattsun/coopy?branch=master)

Under construction

## Installation & Run

```
$ go install github.com/tattsun/coopy
$ COOPY_MYSQL_HOST=localhost:3306 COOPY_MYSQL_USER=coopyuser COOPY_MYSQL_PASSWORD=hogehoge COOPY_MYSQL_DATABASE=coopy coopy
```

## Development

```
$ git clone https://github.com/tattsun/coopy.git
$ cd $GOPATH/src/github.com/tattsun/coopy
$ make install-dev
$ make watch
```
