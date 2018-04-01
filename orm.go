package main

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var engine *xorm.Engine

// InitEngine initializes our Database Connection
func InitEngine() *xorm.Engine {
	var err error
	engine, err = xorm.NewEngine("postgres", "postgres://connService@localhost/learningGo?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	engine.SetMapper(core.GonicMapper{})
	engine.Sync(new(Task))
	return engine
}
