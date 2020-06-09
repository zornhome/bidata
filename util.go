package bidata

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)
const (
	DateTimeFormat = "01-02-15405"
	_out_file = "zorn"
)
var (
	_default_gen_folder = "generate"
	_default_gen_file = "%s/%s-%%s.go"
)
/*
將字串slice寫入檔案
*/
type stringSliceWriter struct {
	// file path
	path string
}
func newStringSliceWriter(fileName string) *stringSliceWriter {
	return &stringSliceWriter{
		fmt.Sprintf( _default_gen_file, _default_gen_folder, fileName),
	}
}

func (w *stringSliceWriter) Write(filename string, lines []string, newline bool) error {
	//加上日期時間做檔名結尾
	fileName := fmt.Sprintf(w.path, time.Now().Format(DateTimeFormat))
	if filename != "" {
		fileName = fmt.Sprintf("%s/%s", _default_gen_folder, filename)
	}

	file,err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	n := "\n"

	if !newline {
		n = ""
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(strings.Join(lines,n))
	writer.Flush()
	fmt.Printf("資料已寫入%s\n", fileName)
	return nil
}

/*
Index 是這個檔案的進入結構(打開程式時,先從至類的index開始看)
通常這類的struct會是 public的
*/
type UtilIndex struct {
	writer *stringSliceWriter
}

func NewUtilIndex() *UtilIndex {
	return &UtilIndex{
		newStringSliceWriter(_out_file),
	}
}
/*
   如何使用:
   	u := bidata.NewUtilIndex()
   	lines := []string {"abckdfghijklmnopqrstuvwxyz","ABCDEFGHIJKLMNOPQRSTUVWXYZ"}
   	u.Writeln(lines) //將資料寫入generate預設資料夾 (參考: newStringSliceWriter )
    lines := []string { "package main\n", "import (\n", `\t"fmt"`,")"}
	u.Write(lines)
  Writeln :會將slice中的元素隔行寫入
  Write : 會將元素全部寫在一行,所以寫入的字串你要自己先處理好
*/
func (u *UtilIndex) Writeln(filename string, lines []string) {
	err := u.writer.Write(filename, lines, true)
	if err != nil {
		panic(err)
	}
}
func (u *UtilIndex) Write(fileName string, lines []string) {
	err := u.writer.Write(fileName, lines, false)
	if err != nil {
		panic(err)
	}
}

func ReadFile(file string) []byte {
	readFile, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return readFile
}