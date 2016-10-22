package godb

import (
	"fmt"
	"os"
	"bufio"
	"github.com/antonholmquist/jason"
)

type GoDB struct {
	path    string
	dataMap map[string]jason.Object
}

func New(fileName string) *GoDB {
	g := &GoDB{path:fileName, dataMap:make(map[string]jason.Object)}
	fmt.Println(g.path)

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	line, e := Readln(r)
	for e == nil {
		//fmt.Println(line)
		player, _ := jason.NewObjectFromBytes(line)
		fmt.Println(player.GetString("name"))
		line, e = Readln(r)
	}
	return g
}
func Readln(r *bufio.Reader) ([]byte, error) {
	var (
		isPrefix bool = true
		err error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return ln, err
}

func InitDb() {
	fmt.Println("initDb")
}

//func ReadLine(filename string) {
//	f, err := os.Open(filename)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer f.Close()
//	r := bufio.NewReader(f)
//	line, isPrefix, err := r.ReadLine()
//	for err == nil && !isPrefix {
//		s := string(line)
//		fmt.Println(s)
//		line, isPrefix, err = r.ReadLine()
//	}
//	if isPrefix {
//		fmt.Println("buffer size to small")
//		return
//	}
//	if err != io.EOF {
//		fmt.Println(err)
//		return
//	}
//}