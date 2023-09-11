package Clear

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearTerminal () {

	var cmd *exec.Cmd;

	if runtime.GOOS == "windowns" {
		cmd = exec.Command("cls");
	} else {
		cmd =  exec.Command("clear");
	}

	cmd.Stdout = os.Stdout;
	cmd.Run()
}