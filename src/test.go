package main

import (
	"os"
	"os/exec"
	"fmt"
	"bytes"
)

func main() {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(`C:\3rdPartySoftware\Git\bin\git.exe`,[]string{"config","--get", "remote.origin.url"}...)
	cmd.Dir = `C:\Users\IBM_ADMIN\GoProjects\src\hub.jazz.net\git\gheorghe\blGo`
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
  	  fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
  	  os.Exit(1)
  	}
  	fmt.Printf("%s\n", out)
}
