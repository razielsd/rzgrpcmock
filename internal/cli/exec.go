package cli

import (
	"fmt"
	"os/exec"
)

func ExecCmd(dirName, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dirName
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

