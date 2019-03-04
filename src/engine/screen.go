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
		fmt.Println("Invalid row index passed in LoadRow(", rowIndex, ")")
		panic(-1)
	}

	d.Details[rowIndex] = row
}
