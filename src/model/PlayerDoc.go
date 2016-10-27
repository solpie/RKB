package model

import (
	"utils/jex"
)

type PlayerDoc  struct {
	jex.JsonEx
}

func (p *PlayerDoc)Init() *PlayerDoc {
	p.Load([]byte(`{
			"id":"",
			"playerNum":"",
			"intro":"",
			"height":"",
			"weight":"",
			"avatar":"",
			"ftName":"",
			"name":""
			}`))
	return p
}