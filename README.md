## [gostruct](#)

Struct Generator from database

gostruct will generate struct from your DB

[![Build Status](https://travis-ci.org/golangid/gostruct.svg?branch=master)](https://travis-ci.org/golangid/gostruct)
[![codecov](https://codecov.io/gh/golangid/gostruct/branch/master/graph/badge.svg)](https://codecov.io/gh/golangid/gostruct)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/golangid/gostruct/blob/master/LICENSE)

## Index

* [Support](#support)
* [Getting Started](#getting-started)
* [Example](#example)

You can file an [Issue](https://github.com/golangid/gostruct/issues/new).

## Support
Currently only support Mysql for now.

## Getting Started

#### Download

```shell
go get -u github.com/golangid/gostruct
```
## Example

#### Creating Configurations
To generate the models, you need to create `config.json` in you project folders.

```json
{
  "type":"mysql",
  "user":"root",
  "password":"password",
  "host":"127.0.0.1",
  "port":"33060",
  "dbname":"article"
}
```
then run from terminal

```bash
$GOPATH/bin/gostruct generate

# Example
~/go/bin/gostruct generate
```

![Example to use Faker](gostruct.gif)
