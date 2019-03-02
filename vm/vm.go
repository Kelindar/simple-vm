package vm

import (
	"errors"
)

// Code represents an instruction code
type code uint32

// Type returns the instruction type
func (c code) Type() uint32 {
	return (0xc0000000 & uint32(c)) >> 30
}

// Data returns the instruction data
func (c code) Data() uint32 {
	return (0x3fffffff & uint32(c))
}

const (
	maxStack = 16
)

var (
	errStackOverflow = errors.New("vm: stack overflow")
)

// VM ...
type VM struct {
	running bool     // Whether the VM is running or not
	stack   []uint32 // The stack
	memory  []uint32 // The Memory
	pc      int32    // Program counter
	sp      int32    // Stack pointer
	tp      uint32   // Type register
	dt      uint32   // Data register
}

// NewVM creates a new virtual machine.
func NewVM(memorySize int) *VM {
	return &VM{
		stack:  make([]uint32, maxStack),
		memory: make([]uint32, memorySize),
		sp:     -1,
		pc:     -1,
	}
}

// Run runs the VM until it's halted or the cycles a.
func (vm *VM) Run(cycles int) error {
	vm.pc = -1
	vm.sp = -1
	vm.running = true
	for cycle := 0; cycle < cycles && vm.running; cycle++ {
		vm.pc++
		instruction := vm.memory[vm.pc]
		vm.tp = code(instruction).Type()
		vm.dt = code(instruction).Data()

		if vm.tp == 1 {
			vm.execute(vm.dt)
			continue
		}

		vm.sp++
		vm.stack[vm.sp] = vm.dt
		if vm.sp == maxStack {
			return errStackOverflow
		}
	}
	return nil
}

// Load loads the program to memory
func (vm *VM) Load(program []uint32) {
	copy(vm.memory, program)
}

// Pop1 pops a value from the stack.
func (vm *VM) Pop1() (v uint32) {
	v = vm.stack[vm.sp]
	vm.sp--
	return
}

// Pop2 pops two values from the stack.
func (vm *VM) Pop2() (v1, v2 uint32) {
	v1 = vm.stack[vm.sp-1]
	v2 = vm.memory[vm.sp]
	vm.sp -= 2
	return
}

// Store stores a value on the stack.
func (vm *VM) Store(value uint32) {
	vm.sp++
	vm.stack[vm.sp] = value
}

func (vm *VM) execute(instruction uint32) {
	switch instruction {
	case 0: // HALT
		vm.running = false

	case 1: // ADD
		a, b := vm.Pop2()
		vm.Store(a + b)

	case 2: // SUB
		a, b := vm.Pop2()
		vm.Store(a - b)

	case 3: // MUL
		a, b := vm.Pop2()
		vm.Store(a * b)

	case 4: // DIV
		a, b := vm.Pop2()
		vm.Store(a / b)
	}
}
