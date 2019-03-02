package main

import (
	"github.com/kelindar/simple-vm/vm"
)

func main() {
	vm := vm.NewVM(1000)
	vm.Load([]uint32{1, 2, 0x40000001, 3, 0x40000001, 0x40000000})
	err := vm.Run(1000)
	if err != nil {
		println(err.Error())
	}
}
