package main

import (
	"fmt"

	"github.io/nohuplus/cmd"
	"github.io/nohuplus/internal/core"
)

func main() {
	if err := core.InitDB(); err != nil {
		fmt.Println(err)
	}
	defer core.CloseDB()
	core.EnsurePaths()
	cmd.Execute()
}
