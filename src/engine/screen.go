package engine

import (
	"fmt"
	"strings"
	//"strconv"
)

type ScreenDetails struct {
	Id string
	Details [DISPLAYWIDTH]string
}

func (sd *ScreenDetails) Print() {
	for idy := 0; idy < len(sd.Details); idy++ {
		fmt.Print(sd.Details[idy])
	}
}

func (d *ScreenDetails) LoadRow(row string, rowIndex int) {
	if rowIndex < 0 || rowIndex > DISPLAYWIDTH {
		panicString := "Invalid row index passed in LoadRow(" + string(rowIndex) + ")"
		fmt.Println(panicString)
		//panic(panicString)
	}

	d.Details[rowIndex] = row
}

func (d *ScreenDetails) ParseAndHandleDetailsString(line string){
	parts := strings.Split(line, DELIMITER)
	d.ParseAndHandleDetailsArray(parts)
}

func (d *ScreenDetails) ParseAndHandleDetailsArray(parts []string){

	if parts == nil || len(parts) <= 1 {
		panic("Invalid parsing for screen details")
	}

	d.Id = parts[0] //The level Id is always the first and included Id

	if len(parts) == 1 { //We have nothing left to do, exit
		return
	}

	for idx := 1; idx < len(parts); idx++ {

		switch strings.ToLower(parts[idx]) {
		case STARTPOSITION:
			fmt.Println("This shouldn't happen...")
		}
		
	}
}