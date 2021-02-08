package test

import (
	"fmt"

	"github.com/docopt/docopt-go"
)
func Testing()(){
	usage := `tool

Usage:
  
  tool -h | --help
 

Options:
  -h --help     Show this screen.`

	  arguments, _ := docopt.ParseDoc(usage)
	  fmt.Println(arguments)

}
