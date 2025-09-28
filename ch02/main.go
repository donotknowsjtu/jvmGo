package main

import (
	"fmt"
	"jvmGo/ch02/classpath"
	"strings"
)

func main() {
	cmd := ParseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		cmd.PrintUsage()
	} else {
		startJVM(cmd)
	}
}
func startJVM(cmd *Cmd) {
	// ch01是将输入的路径整理后直接输出
	//fmt.Printf("classpath:%s class:%s args:%v\n",
	//	cmd.cpOption, cmd.class, cmd.args)

	// ch02
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath: %v class: %v args: %v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("could not find or load main class: %s\n", cmd.class)
		return
	}
	fmt.Printf("class data: %v\n", classData)
}
