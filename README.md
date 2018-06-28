# go-create-shifter-package for Shifter

> This tool is no longer usable with the termination of Shifter Project V1(2018.03).

## Download 

Get binary from here.

- https://github.com/getshifter/go-create-shifter-package/releases

## How to use

https://support.getshifter.io/user-guide/how-to-import-your-wordpress-site-into-shifter

## Development

If you execute from source code yourself to contribute.

### Run localy

```
$ go get github.com/jhoonb/archivex
$ go run main.go
```

### Build and Release

Build single binary for local os.

```
$ go build main.go
```

Build for multi os(macox, linux, windows).

```
$ mkdir build
$ go get github.com/mitchellh/gox
$ gox -os="linux darwin windows" -arch="386 amd64" -output "build/create-shifter-package_{{.OS}}_{{.Arch}}"
```

upload releases to github. (for maintaner information)

```
$ go get github.com/tcnksm/ghr
$ ghr --replace -u getshifter v0.1.1 build/
```
