# go-create-wp-archive for xxx



## Build

```
mkdir build
gox -os="darwin windows" -output "build/{{.Dir}}_{{.OS}}_{{.Arch}}"
```

## Release

```
ghr --replace -u megumiteam v0.1.0 build/
```
