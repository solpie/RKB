package utils

import (
	"math/rand"
	"time"
	"fmt"
	"reflect"
)

/////////////////////  random string
var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1 << letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n - 1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
//////////////////
type Arr struct {
	//_arr []interface{}
	elemType   reflect.Type
	sliceValue reflect.Value
}

//func (arr *Arr)Push(item interface{}) interface{} {
//	fmt.Println(len(arr._arr))
//	var a = append(arr._arr, item)
//	fmt.Println(len(arr._arr))
//	return a
//}
func (self *Arr)Arr() interface{} {
	return self.sliceValue.Interface()
}
func (self *Arr)New(sample interface{}) *Arr {
	value := reflect.ValueOf(sample)
	self.sliceValue = reflect.MakeSlice(value.Type(), 0, 0)
	self.elemType = reflect.TypeOf(sample).Elem()
	return  self
}
func (self *Arr) Push(e interface{}) bool {
	if reflect.TypeOf(e) != self.elemType {
		return false
	}
	self.sliceValue = reflect.Append(self.sliceValue, reflect.ValueOf(e))
	return true
}
func (self *Arr) ElemType() reflect.Type {
	return self.elemType
}

func NewArr() {

}



func Test() {
	var a = Arr{}
	a.New(make([]int,0))
	a.Push(1)
	a.Push(2)
	inst:=a.Arr().([]int)

	var numArr = []int{1,2}
	numArr = append(numArr,3)
	//Arr2(numArr).Push(3)

	//inst:=([]int).(a.Arr())
	fmt.Println("Arr Test():",inst)
	fmt.Println("Arr2 Test():",numArr)
}
