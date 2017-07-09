package main

import (
	"fmt"

	"github.com/pawlo0/udemyGolangTraining/04_scope/01_package_scope/visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
