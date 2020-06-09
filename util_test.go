package bidata

import (
	"bytes"
	"testing"
)

func TestUtilIndex_Writeln(t *testing.T) {
	giveName := "ln"
	give := []string{"package abc","import (","\t\"wxyz\"","\t\"1234\"",")"}
	want := ReadFile("testdata/Writeln")
	ss := &stringSliceWriter{}
	err := ss.Write(giveName, give, true)
	if err != nil {
		t.Fatal(err)
	}
	got := ReadFile("generate/"+ giveName)
	t.Logf("%s",want)
	t.Logf("%s", got)
	if !bytes.EqualFold(got,want) {
		t.Fatal("golden file different with " + giveName)
	}
}

func TestReadFile(t *testing.T) {
	b := ReadFile("generate/ln")
	t.Logf("%s", b)
}