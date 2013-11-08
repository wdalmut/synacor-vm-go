// The VM management
package synacor

import (

)

const mem_size, reg_size = 32768, 8

// The main memory
type Memory struct {
    program     [mem_size]uint16
    register    [reg_size]uint16
    stack       []uint16
}

// Push a new element on the top of the stack
func (memory *Memory)push(value uint16) {
    memory.stack = append(memory.stack, value)
}

// Pop out a value from the stack memory
func (memory *Memory)pop() uint16 {
    value := memory.stack[len(memory.stack)-1]
    memory.stack = memory.stack[0:len(memory.stack)-1]
    return value
}

// If a value is greater than MEM_SIZE return
// the register value
func (memory *Memory)resolve(data uint16) uint16 {
    if data >= mem_size {
        return memory.register[(data % mem_size)]
    }

    return data
}

// Convert an address to a register address in
// case the address is in the register scope
// otherwise the same address is returned
func convert(address uint16) uint16 {
    if address >= mem_size {
        return address % mem_size
    }

    return address
}

// Read from memory
func (memory *Memory)rawRead(address uint16) uint16 {
    if address >= mem_size  {
        return memory.resolve(address)
    } else {
        return memory.program[address]
    }
}

// Load data into the program memory
func (memory *Memory)Load(address uint16, value uint16) {
    if address >= mem_size {
        memory.register[convert(address)] = value
    } else {
        memory.program[address] = value
    }
}

