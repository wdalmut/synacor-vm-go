package synacor

import (
    "testing"
)

// Test bootstrap from file and check that
// info are in little endian
func TestBootstrapFromFile(t *testing.T) {
    vm := new(VM)

    vm.BootstrapFromFile("test.bin")
    memory := vm.program

    if v := memory[0]; v != 256 {
        t.Errorf("Bootstrap from file error: %v, want %v", v, 256)
    }


    if v := memory[1]; v != 512 {
        t.Errorf("Bootstrap from file error: %v, want %v", v, 512)
    }


    if v := memory[2]; v != 768 {
        t.Errorf("Bootstrap from file error: %v, want %v", v, 768)
    }


    if v := memory[3]; v != 0 {
        t.Errorf("Bootstrap from file error: %v, want %v", v, 0)
    }
}

func TestOpOperation(t *testing.T) {
    vm := new(VM)

    vm.op()

    if p := vm.Cpu.ptr; p != 1 {
        t.Errorf("Error op operation %v, want %v", p, 1)
    }
}

func TestSetOperation(t *testing.T) {
    vm := new(VM)

    vm.Load(0, 1)
    vm.Load(1, 32768)
    vm.Load(2, 1)

    vm.op()
    vm.setOp()

    if data := vm.register[0]; data != 1 {
        t.Errorf("Error set operation %v, want %v", data, 1)
    }

    if (vm.Cpu.ptr != 3) {
        t.Errorf("The set operation don't move forward the pointer %v, want %v", vm.Cpu.ptr, 3)
    }
}

func TestPushOp(t *testing.T) {
    vm := new(VM)

    vm.Load(0, 2)
    vm.Load(1, 141)

    vm.op()
    vm.pushOp()

    if data := vm.stack[0]; data != 141 {
        t.Errorf("Error push operation %v, want %v", data, 141)
    }
}

func TestPopOp(t *testing.T) {
    vm := new(VM)

    vm.Load(0, 2)
    vm.Load(1, 141)
    vm.Load(2, 3)
    vm.Load(3, 141)

    vm.op()
    vm.pushOp()
    vm.op()
    vm.popOp()

    if data := vm.program[141]; data != 141 {
        t.Errorf("Error pop operation %v, want %v", data, 141)
    }
}

func TestEqOp(t *testing.T) {
    vm := new(VM)

    vm.Load(0, 4)
    vm.Load(1, 32768)
    vm.Load(2, 3)
    vm.Load(3, 3)

    vm.Load(4, 4)
    vm.Load(5, 32769)
    vm.Load(6, 3)
    vm.Load(7, 5)

    vm.op()
    vm.eqOp()

    if data := vm.register[0]; data != 1 {
        t.Errorf("Error eq operation %v, want %v", data, 1)
    }

    vm.register[1] = 1
    vm.op()
    vm.eqOp()

    if data := vm.register[1]; data != 0 {
        t.Errorf("Error eq operation %v, want %v", data, 0)
    }
}

func TestGtOp(t *testing.T) {
    vm := new(VM)

    vm.Load(0, 5)
    vm.Load(1, 32768)
    vm.Load(2, 3)
    vm.Load(3, 3)

    vm.Load(4, 5)
    vm.Load(5, 32769)
    vm.Load(6, 5)
    vm.Load(7, 4)

    vm.op()
    vm.gtOp()

    if data := vm.register[0]; data != 0 {
        t.Errorf("Error eq operation %v, want %v", data, 0)
    }

    vm.register[1] = 1
    vm.op()
    vm.gtOp()

    if data := vm.register[1]; data != 1 {
        t.Errorf("Error eq operation %v, want %v", data, 1)
    }
}

func TestJumpOp(t *testing.T) {
    vm := new(VM)

    vm.Load(0, 6)
    vm.Load(1, 32)

    vm.op()
    vm.jump()

    if x := vm.Cpu.ptr; x != 32 {
        t.Errorf("Error jump operation %v, want %v", x, 32)
    }
}

func TestJTOp(t *testing.T) {
    vm := new(VM)

    vm.Load(32768, 0)
    vm.Load(32769, 1)

    vm.Load(0, 7)
    vm.Load(1, 32768)
    vm.Load(2, 12421)
    vm.Load(3, 7)
    vm.Load(4, 32769)
    vm.Load(5, 1)

    vm.op()
    vm.jtOp()

    if x := vm.Cpu.ptr; x != 3 {
        t.Errorf("Error jt, jump on zero op %v", x)
    }

    vm.op()
    vm.jtOp()

    if x := vm.Cpu.ptr; x != 1 {
        t.Errorf("Error jt, non jump on non zero op %v", x)
    }
}

func TestJFOp(t *testing.T) {
    vm := new(VM)

    vm.Load(32768, 1)
    vm.Load(32769, 0)

    vm.Load(0, 8)
    vm.Load(1, 32768)
    vm.Load(2, 12421)
    vm.Load(3, 8)
    vm.Load(4, 32769)
    vm.Load(5, 1)

    vm.op()
    vm.jfOp()

    if x := vm.Cpu.ptr; x != 3 {
        t.Errorf("Error jt, jump on non zero op %v", x)
    }

    vm.op()
    vm.jfOp()

    if x := vm.Cpu.ptr; x != 1 {
        t.Errorf("Error jt, non jump on zero op %v", x)
    }
}

func TestAddOp(t *testing.T) {
    vm := new(VM)

    vm.Load(0, 9)
    vm.Load(1, 32768)
    vm.Load(2, 1)
    vm.Load(3, 2)

    vm.op()
    vm.addOp()

    if x := vm.register[0]; x != 3 {
        t.Errorf("Error add op %v, want %v", x, 3)
    }
}

func TestMultOp(t *testing.T) {
    vm := new(VM)

    vm.Load(0, 10)
    vm.Load(1, 32768)
    vm.Load(2, 1)
    vm.Load(3, 2)

    vm.op()
    vm.multOp()

    if x := vm.register[0]; x != 2 {
        t.Errorf("Error mult op %v, want %v", x, 2)
    }
}

func TestModOp(t *testing.T) {
    vm := new(VM)

    vm.Load(0, 11)
    vm.Load(1, 32768)
    vm.Load(2, 5)
    vm.Load(3, 4)

    vm.op()
    vm.modOp()

    if x := vm.register[0]; x != 1 {
        t.Errorf("Error mod op %v, want %v", x, 1)
    }
}

func TestAndOp(t *testing.T) {
    vm := new(VM)

    vm.Load(0, 12)
    vm.Load(1, 32768)
    vm.Load(2, 0x00FF)
    vm.Load(3, 0x0F80)

    vm.op()
    vm.andOp()

    if x := vm.register[0]; x != 0x0080 {
        t.Errorf("Error and op %v, want %v", x, 0x0080)
    }
}

func TestOrOp(t *testing.T) {
    vm := new(VM)

    vm.Load(0, 13)
    vm.Load(1, 32768)
    vm.Load(2, 0x00FF)
    vm.Load(3, 0x0F00)

    vm.op()
    vm.orOp()

    if x := vm.register[0]; x != 0x0FFF {
        t.Errorf("Error and op %v, want %v", x, 0x0FFF)
    }
}

func TestNotOp(t *testing.T) {
    vm := new(VM)

    vm.Load(0, 14)
    vm.Load(1, 32768)
    vm.Load(2, 0x0FFF)

    vm.op()
    vm.notOp()

    if x := vm.register[0]; x != 0x7000 {
        t.Errorf("Error not op %v, want %v", x, 0x7000)
    }
}

func TestCallOp(t *testing.T) {
    vm := new(VM)

    vm.Load(0, 17)
    vm.Load(1, 100)

    vm.op()
    vm.callOp()

    if x := vm.ptr; x != 100 {
        t.Errorf("Error call op, ptr %v, want %v", vm.ptr, 100)
    }
}

func TestRetOp(t *testing.T) {
    vm := new(VM)

    vm.push(100)

    vm.Load(0, 18)

    vm.op()
    vm.retOp()

    if x := vm.ptr; x != 100 {
        t.Errorf("Error ret op, ptr %v, want %v", vm.ptr, 100)
    }
}

