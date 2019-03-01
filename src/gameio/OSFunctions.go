package gameio

import (
	//"fmt"
    "os"
	//"log"
	"os/exec"
    "runtime"
    //"time"
)

var(
	OSListForClear map[string] func() //list of OS supported with corresponding function
)

func init(){
	OSListForClear = make(map[string]func())

	OSListForClear["linux"] = func(){
		cmd := exec.Command("clear") 
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	OSListForClear["windows"] = func(){
		cmd := exec.Command("cmd", "/c", "cls") 
        cmd.Stdout = os.Stdout
        cmd.Run()
	}
}

func ClearConsole(){
	if action, found := OSListForClear[runtime.GOOS]; found {
		action()
	} else{
		panic("You are playing on an unsupported OS. To disable clear functionality pass noclear to app start")
	}
}
