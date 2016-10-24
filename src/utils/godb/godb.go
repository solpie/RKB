package godb

import (
	. "fmt"
	"os"
	"bufio"
	"utils/jex"
	"utils"
)

type GoDB struct {
	_path    string
	_dataMap map[string]*jex.JsonEx
}

func (g *GoDB) Path() string {
	return g._path
}

func (g *GoDB)flush() {
	var docCount = 0
	var data = ""
	for _, value := range g._dataMap {
		//Print(key, ":")
		//value.Value
		var jsonStr = value.String()
		//Println(jsonStr)
		//jsonStr = string(value)
		data += jsonStr + "\n"
		docCount++
	}
	//Println(data)
	if data != "" {
		f, err := os.OpenFile(g._path, os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString(data)

		Println("flush docCount:", docCount)
	}

}
func (g *GoDB)Find() {

}
func (g *GoDB)Update(jo *jex.JsonEx) {
	//var _id = jo.GetString("_id")
	//if _id != nil {
	//	g._dataMap[_id] = jo
	//}
	//
	//if g._dataMap[_id] != nil {
	//	Println("_id exist", _id)
	//	//_id = utils.RandStringBytesMaskImprSrc(7)
	//}
	g.Insert(jo)
}

func (g *GoDB)Insert(jo *jex.JsonEx) {
	var _id string
	_id = utils.RandStringBytesMaskImprSrc(7)
	for {
		if g._dataMap[_id] != nil {
			Println("_id exist", _id)
			_id = utils.RandStringBytesMaskImprSrc(7)
		} else {
			break
		}
	}
	jo.SetP(_id, "_id")
	g._dataMap[_id] = jo

	Println("insert doc:", jo.String())
	f, err := os.OpenFile(g._path, os.O_APPEND, 0666)
	if err != nil {
		Println("error Insert ", err)
	}
	f.WriteString(jo.String())

	defer f.Close()
}

func (g *GoDB) Init(fileName string) {
	f, err := os.OpenFile(fileName, os.O_CREATE, 0666)
	if err != nil {
		Println("error opening file= ", err)
		//os.Exit(1)
	}
	g._path = fileName

	g._dataMap = make(map[string]*jex.JsonEx)
	//g.DataMap = make(map[string]string)
	r := bufio.NewReader(f)
	line, e := readLine(r)
	var docCount = 0
	for e == nil {
		docCount++
		//fmt.Println(line)

		//jsonObj, _ := jason.NewObjectFromBytes(line)
		jsonObj := jex.Load(line)
		var _id = jsonObj.GetString("_id")//Path("_id").Data().(string)
		//var _id, _ = jsonObj.Path("_id").Data().(string)
		//var _id, _ = jsonObj.GetString("_id")
		//Println(_id)
		g._dataMap[_id] = jsonObj
		line, e = readLine(r)
	}
	Println("init db:", g._path, "count:", docCount)

	g.flush()
}

func readLine(r *bufio.Reader) ([]byte, error) {
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
