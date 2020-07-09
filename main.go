package main

import (
	"fmt"
  "github.com/Sergii-Kirichok/structTest/module"
)

type Modules struct {
	Modules 	map[string]*module.Module //string == vendor Name
}

func NewModules() *Modules{
	return &Modules{Modules: make(map[string]*module.Module)}
}


func (mS *Modules) SMProcessing() {
	name:="Vname"
	m := module.NewModule()

	m.SetVendorName(name)
	m.SetStatusCur("Running")
	m.SetStatusAsk("Stop")
	mS.Modules[name] = m

	fmt.Println(mS)
	fmt.Println(mS.Modules[name])
}

func main() {
	s:=NewModules()
	s.SMProcessing()
}