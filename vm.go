package main

import (
    "fmt"
    "github.com/wdalmut/synacor-vm-go/synacor"
    "flag"
)

func main() {
    flag.Parse()
    binPath := flag.Arg(0);

    fmt.Printf("%v", binPath)

    vm := new(synacor.VM)
    vm.BootstrapFromFile(binPath)

    vm.Run()
}


