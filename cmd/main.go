package main

import (
	"bidata"
	"time"
)

func main(){
	u := bidata.NewUtilIndex()
	formatlines :=  []string{"package abc","import (","\t\"wxyz\"","\t\"1234\"",")"}
	u.Writeln("Writeln", formatlines)
	time.Sleep(time.Second * 2)
/*	lines := []string {
		"package main","import (","\"fmt\"","\"strings\"",")",
	}
	u.Writeln("ln", lines)
*/
}




























































































