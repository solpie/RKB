package model

import (
	"utils/godb"
	//"fmt"
	"sync"
)


type DbModel struct {
	O         interface{};
	PlayerDb  *godb.GoDB
	PlayerMap map[string]*PlayerDoc
}

var inst *DbModel
var once sync.Once

func Db() *DbModel {
	once.Do(func() {
		inst = &DbModel{}
		inst.PlayerDb = godb.Load("./db/player.db")
		inst.PlayerMap = make(map[string]*PlayerDoc)
	})
	return inst
}