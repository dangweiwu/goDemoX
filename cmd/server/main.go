package main

import (
	"DEMOX_ADMINAUTH/option"
	"github.com/jessevdk/go-flags"
)

// @base| xx系统管理 | v0.0.1
// @desc| 系统管理 2024年 12月 1日
func main() {
	p := flags.NewParser(&option.Opt, flags.Default)
	p.ShortDescription = "v1.0 server"
	p.LongDescription = `v1.0 Server`
	p.Parse()
}
