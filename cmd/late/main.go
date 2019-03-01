package main

import "github.com/blang/late/pkg/late/cmd"

var SemVer string = "0.0.0-unknown"

func main() {
	cmd.SemVer = SemVer
	cmd.Execute()
}
