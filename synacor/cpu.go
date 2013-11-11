package synacor

import (
    "io/ioutil"
    "fmt"
    "os"
)

const modulo = 32768

// The cpu (only ptr)
type Cpu struct {
    ptr         uint16
}

// The VM is a cpu and a memory
type VM struct {
   Cpu
   Memory
}

// Reset the program counter to 0
func (cpu *Cpu)reset() uint16 {
    cpu.ptr = 0
    return cpu.ptr
}

// Retrive the actual program counter
// and after increase it
func (cpu *Cpu)inc() uint16 {
    ptr := cpu.ptr
    cpu.ptr += 1
    return ptr
}

func (vm *VM)BootstrapFromFile(path string) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        panic("Unable to read your file")
        return
    }
    n, j := len(data), 0
    for i := 0; j < n; i++ {
        vm.Load(uint16(i), uint16(data[j]) | (uint16(data[j+1]) << 8))
        j = j + 2
    }
}

// Read a new word from the memory
func (vm *VM)op() uint16 {
    return vm.rawRead(vm.Cpu.inc())
}

// Run the program loaded into the main
// memory
func (vm *VM)Run() {
    for {
        word := vm.op()
        switch word {
            case 0:
                os.Exit(0)
            case 1:
                vm.setOp()
            case 2:
                vm.pushOp()
            case 3:
                vm.popOp()
            case 4:
                vm.eqOp()
            case 5:
                vm.gtOp()
            case 6:
                vm.jump()
            case 7:
                vm.jtOp()
            case 8:
                vm.jfOp()
            case 9:
                vm.addOp()
            case 10:
                vm.multOp()
            case 11:
                vm.modOp()
            case 12:
                vm.andOp()
            case 13:
                vm.orOp()
            case 14:
                vm.notOp()
            case 15:
                vm.rmemOp()
            case 16:
                vm.wmemOp()
            case 17:
                vm.callOp()
            case 18:
                vm.retOp()
            case 19:
                vm.outOp()
            case 20:
                vm.inOp()
            case 21:
                vm.noopOp()
        }
    }
}

// Set a value into a register
func (vm *VM)setOp() {
    a := vm.op()
    b := vm.resolve(vm.op())

    vm.Load(a, b)
}

// Push a new value onto the stack
func (vm *VM)pushOp() {
    a := vm.resolve(vm.op())
    vm.push(a)
}

// Pop a value from the stack
func (vm *VM)popOp() {
    a := vm.op()
    value := vm.pop()

    vm.Load(a, value)
}

// Check equals args
func (vm *VM)eqOp() {
    a := vm.op()
    b := vm.resolve(vm.op())
    c := vm.resolve(vm.op())

    vm.Load(a, 0)
    if (b == c) {
        vm.Load(a, 1)
    }
}

// Check greater than
func (vm *VM)gtOp() {
    a := vm.op()
    b := vm.resolve(vm.op())
    c := vm.resolve(vm.op())

    vm.Load(a, 0)
    if (b > c) {
        vm.Load(a, 1)
    }
}

// Jump to another location
func (vm *VM)jump() {
    a := vm.op()
    vm.Cpu.ptr = a
}

// Jump if not zero
func (vm *VM)jtOp() {
    a := vm.resolve(vm.op())
    b := vm.resolve(vm.op())

    if a != 0 {
        vm.Cpu.ptr = b
    }
}

// Jump on zero
func (vm *VM)jfOp() {
    a := vm.resolve(vm.op())
    b := vm.resolve(vm.op())

    if a == 0 {
        vm.Cpu.ptr = b
    }
}

// Add 2 figures
func (vm *VM)addOp() {
    a := vm.op()
    b := vm.resolve(vm.op())
    c := vm.resolve(vm.op())

    vm.Load(a, (b + c) % modulo)
}

// Multiply 2 figures
func (vm *VM)multOp() {
    a := vm.op()
    b := vm.resolve(vm.op())
    c := vm.resolve(vm.op())

    vm.Load(a, (b * c) % modulo)
}

func (vm *VM)modOp() {
    a := vm.op()
    b := vm.resolve(vm.op())
    c := vm.resolve(vm.op())

    vm.Load(a, b % c)
}

func (vm *VM)andOp() {
    a := vm.op()
    b := vm.resolve(vm.op())
    c := vm.resolve(vm.op())

    vm.Load(a, b & c)
}

func (vm *VM)orOp() {
    a := vm.op()
    b := vm.resolve(vm.op())
    c := vm.resolve(vm.op())

    vm.Load(a, b | c)
}

func (vm *VM)notOp() {
    a := vm.op()
    b := vm.resolve(vm.op())

    vm.Load(a, (^b) & 0x7FFF)
}

func (vm *VM)rmemOp() {
    a := vm.op()
    b := vm.rawRead(vm.resolve(vm.op()))

    vm.Load(a, b)
}

func (vm *VM)wmemOp() {
    a := vm.resolve(vm.op())
    b := vm.resolve(vm.op())

    vm.Load(a, b)
}

func (vm *VM)callOp() {
    a := vm.resolve(vm.op())

    vm.push(vm.ptr)
    vm.ptr = a
}

func (vm *VM)retOp() {
    vm.ptr = vm.pop()
}

func (vm *VM)outOp() {
    a := vm.resolve(vm.op())

    fmt.Printf("%c", a)
}

func (vm *VM)inOp() {
    a := vm.op()

    buf := make([]byte, 1)
    os.Stdin.Read(buf)
    vm.Load(a, uint16(buf[0]))
}

func (vm *VM)noopOp() {

}

