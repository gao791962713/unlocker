// SPDX-FileCopyrightText: © 2014-2022 David Parsons
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"os"

	"github.com/drdonk/golocker/vmwpatch"
)

func main() {
	vmwpatch.IsAdmin()

	// Print titles
	fmt.Printf("DumpSMC %s\n", vmwpatch.VERSION)
	fmt.Printf("==============\n")
	fmt.Printf("%s \n\n", vmwpatch.COPYRIGHT)
	// Get and check file passed as parameter
	var filename string
	if len(os.Args) < 2 {
		fmt.Printf("Please pass a file name!\n")
		return
	} else {
		filename = os.Args[1]
	}

	// Read the file
	contents := vmwpatch.LoadFile(filename)
	vmwpatch.DumpSMC(contents)
	patchFlag, patchStatus, hash256 := vmwpatch.IsSMCPatched(contents)
	fmt.Printf("\nPatch Status: %s (%d)\nSHA256: %s\n", patchStatus, patchFlag, hash256)

}
