package godb

import (
	. "fmt"
	"utils/jex"
	"testing"
	"path/filepath"
	"os"
)

func Test(t *testing.T) {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		Println(err)

	}
	Println(dir)
	//var playerDb = new(GoDB)
	//playerDb.Init("./player2.db")

	var newDb = new(GoDB)
	newDb.Init("./test.db")

	//playerDb.Insert(nil)
	//Println(playerDb.Path())

	Println("test insert")

	//jex :=jex.JsonEx{jsonObj}
	var jo = jex.Load([]byte(`{
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
	newDb.Insert(jo)
	Println(jo.GetNumber("outter.inner.value1"))
	jo.SetP(998, "outter.inner.value2")

	jo2 := jo.Clone()
	jo2.SetP(9005, "outter.inner.value2")

	//Println(jo.String())
	//Println(jo2.String())

	////test sort

	var jexArr = []*jex.JsonEx{}
	jexArr = append(jexArr, jo)
	jexArr = append(jexArr, jo2)
	Println(jexArr)


	jexArr =jex.Sort(jexArr,"+outter.inner.value2")
	jexArr =jex.Sort(jexArr,"-outter.inner.value2")
	//s := NewSorter().ByKeys([]string{
	//	"-score",
	//})
	//var score = "score"
	//data := []map[string]interface{}{
	//	{"value": jo, score: jo.GetNumber("outter.inner.value2")},
	//	{"value": jo2, score: jo.GetNumber("outter.inner.value2")},
	//}
	//
	//s.Sort(data)
	////Println(data)
	//
	//var sortJexArr = make([]*jex.JsonEx, len(jexArr))
	//for i := range data {
	//	sortJexArr[i] = data[i]["value"].(*jex.JsonEx)
	//}
	//Println(jexArr)
}