package main

import (
	"myApp/packageone"
)

var myVar = "myVar variable"

func main() {

	var blockVar = "blockVar variable"

	packageone.PrintMe(myVar, blockVar, packageone.PackageVar)

}
