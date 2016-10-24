package jex

import (
	"github.com/Jeffail/gabs"
	"fmt"
	"github.com/coryb/sorty"
)

type JsonEx struct {
	ctn *gabs.Container
}

func Type() interface{} {
	var t JsonEx
	return t
}

func (jex *JsonEx)Ctn(path string) *gabs.Container {
	return jex.ctn
}
func (jex *JsonEx)GetArray(path string) []*JsonEx {
	var children, _ = jex.ctn.Path(path).Children()
	var a = make([]*JsonEx, len(children))
	for i := 0; i < len(children); i++ {
		a[i] = Load(children[i])//{array[i]}
	}
	return a
}

func (jex *JsonEx)String() string {
	return jex.ctn.String()
}
func (jex *JsonEx)Data() interface{} {
	return jex.ctn.Data()
}

func (jex *JsonEx)SetP(value interface{}, path string) (*JsonEx, error) {
	_, err := jex.ctn.SetP(value, path)
	return jex, err
}

func (jex *JsonEx)GetString(path string) string {
	v, _ := jex.ctn.Path(path).Data().(string)
	return v
}

func (jex *JsonEx)GetNumber(path string) float64 {
	v, _ := jex.ctn.Path(path).Data().(float64)
	return v
}

func Sort(jex []*JsonEx, opPath string) []*JsonEx {
	//fmt.Println(string(opPath[0]))
	op := string(opPath[0])
	path := string(opPath[1:len(opPath) - 1])
	s := sorty.NewSorter().ByKeys([]string{
		op + "key",
	})
	var tmpArr = make([]map[string]interface{}, len(jex))
	for i := range jex {
		tmpArr[i] = map[string]interface{}{"jex": jex[i], "key": jex[i].GetNumber(path)}
	}
	s.Sort(tmpArr)
	//
	var sortJexArr = make([]*JsonEx, len(jex))
	for i := range tmpArr {
		sortJexArr[i] = tmpArr[i]["jex"].(*JsonEx)
	}
	//fmt.Println(sortJexArr)
	return sortJexArr
}

func (jex *JsonEx)Clone() *JsonEx {
	return Load(jex.ctn.Bytes())
}

func (jex *JsonEx)Load(param  interface{}) *JsonEx {
	switch inst := param.(type){
	case []byte:
		ctn, _ := gabs.ParseJSON(inst)
		jex.ctn = ctn;
	case *gabs.Container:
		jex.ctn = inst;
	default:
		fmt.Println("unknow")
	}
	return jex
}
func Load(param  interface{}) *JsonEx {
	return new(JsonEx).Load(param)
}