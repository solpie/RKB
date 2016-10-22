package godb

import (
	"fmt"
	"os"
	"bufio"
	"github.com/antonholmquist/jason"
)

type GoDB struct {
	Path    string
	DataMap map[string]*jason.Object
}

func (g *GoDB) Init(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}
	g.DataMap = make(map[string]*jason.Object)
	r := bufio.NewReader(f)
	line, e := Readln(r)
	for e == nil {
		//fmt.Println(line)
		jsonObj, _ := jason.NewObjectFromBytes(line)
		var _id, _ = jsonObj.GetString("id")
		g.DataMap[_id] = jsonObj
		//fmt.Println(player.GetString("name"))
		line, e = Readln(r)
	}
	fmt.Println(g.DataMap["119"])
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