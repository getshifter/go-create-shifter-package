# go-create-shifter-package for Shifter

Get binary from here. https://github.com/getshifter/go-create-shifter-package/releases


## development

### Build

```
mkdir build
gox -os="darwin windows" -arch="386 amd64" -output "build/create-shifter-package_{{.OS}}_{{.Arch}}"
```

### Release

```
ghr --replace -u megumiteam v0.1.1 build/
```
