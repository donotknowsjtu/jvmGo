package main

import (
	"fmt"
	"jvmGo/ch01/cmd"
)

func main() {
	cmd := cmd.ParseCmd()
	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		cmd.PrintUsage()
	} else {
		startJVM(cmd)
	}
}
func startJVM(cmd *cmd.Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.CpOption, cmd.Class, cmd.Args)
}
