package main

import say_cgo "plugins/say-cgo"

type sayPlugin string

func (p sayPlugin) Say(name, something string) {
	say_cgo.Say(name, something)
}

// SayPlugin export symbol
var SayPlugin sayPlugin
