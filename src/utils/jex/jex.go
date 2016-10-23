package jex

import (
	"github.com/Jeffail/gabs"
	"fmt"
)

type JsonEx struct {
	ctn *gabs.Container
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