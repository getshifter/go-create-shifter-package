# go-create-shifter-package for Shifter



## Build

```
mkdir build
gox -os="darwin windows" -output "build/create-shifter-package_{{.OS}}_{{.Arch}}"
```

## Release

```
ghr --replace -u megumiteam v0.1.0 build/
```
