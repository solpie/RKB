package godb

import (
	. "fmt"
	"os"
	"bufio"
	"github.com/Jeffail/gabs"
)

type GoDB struct {
	_path    string
	_dataMap map[string]*gabs.Container
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
func (g *GoDB)Insert(jo *gabs.Container) {
	var _id = ""
	// _id2,_:= jo.GetString("_id")
	//if _id2!=nil{
	//	_id = _id2
	//}
	_id = RandStringBytesMaskImprSrc(7)
	for {
		if g._dataMap[_id] != nil {
			Println("_id exist", _id)
			_id = RandStringBytesMaskImprSrc(7)
		} else {
			break
		}
	}
	Println("new _id", _id)
	jo.SetP(_id,"_id")

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

	g._dataMap = make(map[string]*gabs.Container)
	//g.DataMap = make(map[string]string)
	r := bufio.NewReader(f)
	line, e := readLine(r)
	var docCount = 0
	for e == nil {
		docCount++
		//fmt.Println(line)

		//jsonObj, _ := jason.NewObjectFromBytes(line)
		jsonObj, _ := gabs.ParseJSON(line)
		var _id, _ = jsonObj.Path("_id").Data().(string)
		//var _id, _ = jsonObj.Path("_id").Data().(string)
		//var _id, _ = jsonObj.GetString("_id")
		//Println(_id)
		g._dataMap[_id] = jsonObj
		line, e = readLine(r)
	}
	Println("init db:", g._path, "count:", docCount)

	g.flush()
	//g.Insert(nil)
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

func Test() {
	var playerDb = new(GoDB)
	playerDb.Init("./db/player2.db")

	var newDb = new(GoDB)
	newDb.Init("./db/ft.db")

	//playerDb.Insert(nil)
	Println(playerDb.Path())

	jsonObj, _ := gabs.ParseJSON([]byte(`{
    "outter":{
        "inner":{
            "value1":10,
            "value2":22
        },
        "alsoInner":{
            "value1":20
        }
    }
}`))
	Println("test insert")
	newDb.Insert(jsonObj)


	//var value float64
	//var ok bool

	var value, ok = jsonObj.Path("outter.inner.value1").Data().(float64)
	Println(value, ok)
	//jsonObj := gabs.New()
	// or gabs.Consume(jsonObject) to work on an existing map[string]interface{}

	//jsonObj.Set(99, "outter", "inner", "value")
	jsonObj.SetP(20, "outter.inner.value2")
	//jsonObj.Set(30, "outter", "inner2", "value3")

	Println(jsonObj.String())

}