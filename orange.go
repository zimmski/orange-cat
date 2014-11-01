package main

import (
	"fmt"
	"strconv"
)

func NewOrange(filepath string) *Orange {
	return &Orange{filepath, false}
}

type Orange struct {
	filepath string
	useBasic bool
}

func (o *Orange) UseBasic() {
	o.useBasic = true
}

func (o *Orange) Run(port int) {
	portString := ":" + strconv.Itoa(port)
	fmt.Println(portString)
}
