package controller

import (
	"github.com/claudioed/podkeeper/pkg/controller/podkeeper"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, podkeeper.Add)
}
