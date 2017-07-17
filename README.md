# go-create-shifter-package for Shifter

## Download 

Get binary from here.

- https://github.com/getshifter/go-create-shifter-package/releases


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
$ github.com/mitchellh/gox
$ gox -os="darwin windows" -arch="386 amd64" -output "build/create-shifter-package_{{.OS}}_{{.Arch}}"
```

upload releases to github. (for maintaner information)

```
$ ghr --replace -u megumiteam v0.1.1 build/
```
