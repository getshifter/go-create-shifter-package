package main

import (
	"fmt"
	"os"
	"log"
	"github.com/jhoonb/archivex"
	"path/filepath"
)

func showUsageAndExit ()  {
	var stdin string
	usage := `Showing Help message...

==> Put files below and please retry.

=============================================
./
├── go-create-wp-archive[.exe]
├── webroot/
│   ├── wp-admin/
│   ├── wp-content/
│   ├── wp-includes/
│   ├── (All WordPress Files......)
│   ├── index.php
│   ├── wp-activate.php
│   ├── wp-blog-header.php
│   ├── wp-comments-post.php
│   ├── wp-config.php
│   └── xmlrpc.php
└── wp.sql (dumped from mysql)
=============================================`
	log.Println(usage)
	fmt.Scanln(&stdin)
	os.Exit(1)
}

var fpaths []string
var checkpath []string

func main () {
	dir := filepath.Dir(os.Args[0])
	fmt.Println(dir)
	os.Chdir(dir)

	checkStructure()
	createArchive()
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func checkStructure(){
	checkpath = append(checkpath, "wp.sql")
	checkpath = append(checkpath, "webroot/index.php")
	checkpath = append(checkpath, "webroot/wp-admin/admin.php")
	checkpath = append(checkpath, "webroot/wp-content/index.php")
	checkpath = append(checkpath, "webroot/wp-config.php")

	for _, fpath := range checkpath {
		if Exists(fpath) {
			// any message ?
		} else {
			log.Println("Missing Some Files !!")
			showUsageAndExit()
		}
	}
}

func createArchive() {
	arcive_name := "wordpress.zip"
	zip := new(archivex.ZipFile)
	if Exists(arcive_name) {
		os.Remove(arcive_name)
	}

	zip.Create(arcive_name)
	zip.AddFile("wp.sql")
	zip.AddAll("webroot", true)
	zip.Close()
}

