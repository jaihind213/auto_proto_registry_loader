package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const pkg = "pocket_registry"

//note: this can be implemented as bash script too! since we are in golang, why not just do in golang its easy!
//scan through the protobuff generated golang code and prepare for loading into protobuff registry object.
func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		panic("require 2 arguments. 1st argument: <directory full path containing protobuff generated golang code>, 2nd argument: <golang module url>. example usage: go run run -mod=mod github.com/jaihind213/auto_proto_registry_loader/load/ /home/project/proto/golang/ github.com/dc/batman")
	}

	goProtoCodeDir := args[0]
	modUrl := args[1]
	projectDirName := modUrl[strings.LastIndex(modUrl, "/")+1:]

	_ = os.Mkdir(goProtoCodeDir+"/"+pkg, 0755)
	fyle, fErr := os.OpenFile(goProtoCodeDir+"/"+pkg+"/register_proto_defs.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if fErr != nil {
		fmt.Println("failed to create file register_proto_defs")
		log.Fatal(fErr)
	}
	defer fyle.Close()

	datawriter := bufio.NewWriter(fyle)

	//write lic
	_, _ = datawriter.WriteString("//This file is auto-generated by 'github.com/jaihind213/auto_proto_registry_loader' . please dont modify the file!\" \n")
	//write package line
	_, _ = datawriter.WriteString("package " + pkg + "\n\n")

	if flushErr := datawriter.Flush(); flushErr != nil {
		fmt.Println("failed to auto-generate file register_proto.go ")
		panic(flushErr.Error())
	}

	entries := make(map[string]string, 0)

	walkErr := filepath.Walk(goProtoCodeDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Contains(path, ".pb.go") {
				pkgPath := path[strings.LastIndex(path, projectDirName)+len(projectDirName):]
				lastSlash := strings.LastIndex(pkgPath, string(filepath.Separator))
				pkgPath = pkgPath[:lastSlash]
				importLine := "import _ \"" + modUrl + pkgPath + "\""
				if _, ok := entries[importLine]; !ok {
					entries[importLine] = "-"
					_, _ = datawriter.WriteString(importLine + "\n")
					fmt.Println(importLine)
					if flushErr := datawriter.Flush(); flushErr != nil {
						fmt.Println("failed to auto-generate file register_proto_defs.go ")
						panic(flushErr.Error())
					}
				}
			}

			return nil
		})

	if walkErr != nil {
		fmt.Println("failed to list all files for auto-registration of protos")
		panic(walkErr.Error())
	}

}
