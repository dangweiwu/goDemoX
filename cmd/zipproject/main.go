package main

import (
	"fmt"
	"github.com/dangweiwu/ginpro/pkg"
)

var ignorfile = []string{"zipproject.go", "go.mod", "go.sum", "DEMOX_ADMINAUTH.zip", ".idea", "cover.html", "cover.out", "report.xml", "cmd/zipproject", "cmd/zipproject/main.go"}

func main() {
	if err := pkg.ZipFile("D:/study/gin-template/adminauth/DEMOX_ADMINAUTH", "DEMOX_ADMINAUTH.zip", ignorfile); err != nil {
		fmt.Println("[error]", err)
	}

}
