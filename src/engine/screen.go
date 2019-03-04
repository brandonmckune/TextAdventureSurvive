package engine

import (
	"fmt"
)

type ScreenDetails struct {
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
