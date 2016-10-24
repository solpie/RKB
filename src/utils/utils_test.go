package utils

import (
	"testing"
	"fmt"
	"reflect"
	"utils/jex"
)

type MyInt struct {
	V int
}

func TypeMyType() interface{} {
	var t MyInt
	return t
}
func Test(t1 *testing.T) {
	//reflect.MakeSlice(reflect.SliceOf(myType), 0, 0).Interface()
	var a = Arr{}
	a.New(0)
	a.Push(1)
	a.Push(2)
	inst := a.Arr().([]int)

	//var numArr = []int{1, 2}
	//numArr = append(numArr, 3)
	//Arr2(numArr).Push(3)

	//inst:=([]int).(a.Arr())
	fmt.Println("Arr Test():", inst, len(inst))
	//fmt.Println("Arr2 Test():", numArr)

	var b = NewArr(TypeMyType())

	var ii = MyInt{1}
	b.Push(ii)

	var inst2 = b.Arr().([]MyInt)
	fmt.Println("Arr Test2():", ii.V, inst2)

	if res, ok := b.Arr().([]MyInt); ok {
		fmt.Println(res)
	} else {
		fmt.Println("You are out of luck :(")
	}

	var x = 123
	t := reflect.TypeOf(x)
	fmt.Println(t)



	var a3 = NewArr(jex.Type())
	var jo = jex.Load([]byte(`{"data":0}`))
	a3.Push(jo)
	jo.SetP(1,"data")
	jexArr:=a3.Arr().([]jex.JsonEx)
	fmt.Println("jex Arr:",len(jexArr),a3.Arr().([]jex.JsonEx))
	//st := reflect.SliceOf(t)
	//fmt.Println(st)
	////sd := reflect.MakeSlice(st, 0, 10)
	//sd := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(1)), 0, 0)
	//
	//fmt.Println(sd)
	//tv := reflect.ValueOf(x)
	//sd = reflect.Append(sd, tv, tv, tv)
	//fmt.Println(sd)
	//if res, ok := sd.Interface().([]int); ok {
	//	fmt.Println(res[0])
	//} else {
	//	fmt.Println("You are out of luck :(")
	//}

	//slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(1)), 0, 0).Interface()
	//slice = reflect.Append(slice, reflect.ValueOf(2))
	//
	//fmt.Println(slice)
}

