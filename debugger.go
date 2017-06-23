package debugger

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"syscall"
)

type Breakpoint struct {
	Pid  int
	Addr uintptr
	Orig int
}

func BreakPoint(pid int, addr uintptr) *Breakpoint {
	return &Breakpoint{
		Pid:  pid,
		Addr: addr,
	}
}

func (b *Breakpoint) Enable() {
	origdata := []byte{}
	syscall.PtracePeekText(b.Pid, b.Addr, origdata)
	bBuf := bytes.NewBuffer(origdata)
	binary.Read(bBuf, binary.BigEndian, &b.Orig)
	fmt.Println(b.Orig)

	var n = (b.Orig & 0xFFFFFF00) | 0xCC
	nBuf := bytes.NewBuffer([]byte{})
	binary.Write(nBuf, binary.BigEndian, n)
	fmt.Println(nBuf.Bytes())
	syscall.PtracePokeText(b.Pid, b.Addr, nBuf.Bytes())
}

func (b *Breakpoint) Disable() {
	origdata := []byte{}
	syscall.PtracePeekText(b.Pid, b.Addr, origdata)
	bBuf := bytes.NewBuffer(origdata)
	var x int
	binary.Read(bBuf, binary.BigEndian, &x)
	fmt.Println(x)
	var y = (x & 0xFFFFFF00) | (b.Orig & 0xFF)
	nBuf := bytes.NewBuffer([]byte{})
	binary.Write(nBuf, binary.BigEndian, y)
	fmt.Println(nBuf.Bytes())
	syscall.PtracePokeText(b.Pid, b.Addr, nBuf.Bytes())
}
