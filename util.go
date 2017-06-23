package debugger

import (
	"bytes"
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

var (
	ErrNotExist = errors.New("pid does not exist")
)

// find pid by command "ps aux"
func FindPidByPs(procname string) (int, error) {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return 0, err
	}
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			return 0, err
		}
		if strings.Contains(line, procname) {
			fields := strings.Split(line, " ")
			fiter := make([]string, 0)
			for _, v := range fields {
				if v != "" && v != "\t" {
					fiter = append(fiter, v)
				}
			}
			pid, err := strconv.Atoi(fiter[1])
			if err != nil {
				return 0, err
			}
			return pid, nil
		} else {
			continue
		}
	}
	return 0, ErrNotExist
}
