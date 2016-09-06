package main

import (
	"fmt"
	"github.com/jhoonb/archivex"
	"log"
	"os"
	"path/filepath"
)

func showUsageAndExit() {
	var stdin string
	usage := `Showing Help message...

==> Put files below and please retry.

=============================================
./
├── go-create-shifter-package[.exe]
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

func main() {
	dir := filepath.Dir(os.Args[0])
	log.Println("Working in " + dir)
	os.Chdir(dir)

	checkStructure()
	createArchive()
}

// Exists checks file existence
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func checkStructure() {
	checkpath = append(checkpath, "wp.sql")
	checkpath = append(checkpath, filepath.Join("webroot", "index.php"))
	checkpath = append(checkpath, filepath.Join("webroot", "wp-admin", "admin.php"))
	checkpath = append(checkpath, filepath.Join("webroot", "wp-content", "index.php"))
	checkpath = append(checkpath, filepath.Join("webroot", "wp-config.php"))

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
	arciveName := "wordpress.zip"
	zip := new(archivex.ZipFile)
	if Exists(arciveName) {
		os.Remove(arciveName)
	}

	zip.Create(arciveName)
	zip.AddFile("wp.sql")
	zip.AddAll("webroot", true)
	zip.Close()
}
