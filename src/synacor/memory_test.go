package synacor

import (
    "testing"
)

// Test not necessary resolve operation
func TestNoResolve(t *testing.T) {
    const in, out = 1, 1

    var memory = new(Memory)

    if x := memory.resolve(in); x != out {
        t.Errorf("Resolve fail conversion %v -> %v, want %v", in, x, out)
    }
}

// Test necessary resolve operation
func TestResolve(t *testing.T) {
    const in, out = 32768, 123

    var memory = new(Memory)
    memory.register[0] = 123

    if x := memory.resolve(in); x != out {
        t.Errorf("Resolve fail conversion %v -> %v, want %v", in, x, out)
    }
}

// Test base push and pop
func TestPushAndPopSingle(t *testing.T) {
    memory := new(Memory)
    memory.push(10)
    if data := memory.pop(); data != 10 {
        t.Errorf("Error pop value %v, want %v", data, 10)
    }
}

// Test that values are popped from the top...
func TestMultiplePushAndPop(t *testing.T) {
    memory := new(Memory)
    memory.push(1);
    memory.push(2);
    memory.push(3);

    if data := memory.pop(); data != 3 {
        t.Errorf("Error pop value %v, want %v", data, 3)
    }

    if data := memory.pop(); data != 2 {
        t.Errorf("Error pop value %v, want %v", data, 2)
    }

    if data := memory.pop(); data != 1 {
        t.Errorf("Error pop value %v, want %v", data, 1)
    }
}

func TestLoadAndFetchProgramMemory(t *testing.T) {
    memory := new(Memory)

    memory.Load(0, 1)
    memory.Load(1, 2)
    memory.Load(32768, 1)

    if data := memory.rawRead(0); data != 1 {
        t.Errorf("Error fetch operation %v, want %v", data, 1)
    }

    if data := memory.rawRead(1); data != 2 {
        t.Errorf("Error fetch operation %v, want %v", data, 2)
    }

    if data := memory.register[0]; data != 1 {
        t.Errorf("Error fetch operation %v, want %v", data, 1)
    }
}

func TestRawRead(t *testing.T) {
    memory := new(Memory)

    memory.Load(0, 1)
    memory.Load(1, 32768)
    memory.Load(32768, 1)

    if x := memory.rawRead(0); x != 1 {
        t.Errorf("Error raw read %v, want %v", x, 1)
    }

    if x := memory.rawRead(1); x != 32768 {
        t.Errorf("Error raw read %v, want %v", x, 32768)
    }

    if x := memory.rawRead(32768); x != 1 {
        t.Errorf("Error raw read %v, want %v", x, 1)
    }
}

func TestConvertAddress(t *testing.T) {
    if x:= convert(32768); x != 0 {
        t.Errorf("Error convertion from %v to %v, obtained %v", 32768, 0, x)
    }

    if x:= convert(32769); x != 1 {
        t.Errorf("Error convertion from %v to %v, obtained %v", 32769, 1, x)
    }

    if x:= convert(12); x != 12 {
        t.Errorf("Error convertion from %v to %v, obtained %v", 12, 12, x)
    }
}

