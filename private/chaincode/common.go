package main

import (
	"github.com/s7techlab/cckit/router"
)

//IdExist ... Check if the id exist in world state
func IdExist(id string, context router.Context) bool {
	exits, _ := context.State().Exists(id)
	return exits
}
