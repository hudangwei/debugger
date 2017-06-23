package debugger

import "fmt"

func Example() {
	pid, err := FindPidByPs("/home/oneu/ccpos/dist/ccpos")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pid)

	bp := BreakPoint(pid, 0x014d7da8)
	bp.Enable()
}
