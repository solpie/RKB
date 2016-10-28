package model


type ServerModel struct {
	O         interface{};
	OpUrlMap map[string]string
}

var instServerModel *ServerModel

func SrvModel() *ServerModel {
	once.Do(func() {
		instServerModel = &ServerModel{}
		instServerModel.OpUrlMap = make(map[string]string)
	})
	return instServerModel
}
