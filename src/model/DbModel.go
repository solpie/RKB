package model

import (
	"utils/godb"
	//"fmt"
	"sync"
)

//var PlayerDb *godb.GoDB

//func InitDb() {
//	var PlayerDb = godb.Load("./db/player.db")
//	fmt.Println(PlayerDb.Path())
//}


//package singleton

//import "sync"

type single struct {
	O interface{};
	PlayerDb *godb.GoDB
}
var instantiated *single
var once sync.Once

func Db() *single {
	once.Do(func() {
		instantiated = &single{}
		instantiated.PlayerDb =  godb.Load("./db/player.db")
	})
	return instantiated
}