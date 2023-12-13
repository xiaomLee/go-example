package main

import (
	"log"
	"plugin"
)

type PluginSay interface {
	Say(name, something string)
}

func main() {
	path := "output/sayPlugin.so"
	p, err := plugin.Open(path)
	if err != nil {
		log.Fatalf("error opening plugin: %v path:%s", err, "")
	}
	symbol, err := p.Lookup("SayPlugin")
	if err != nil {
		log.Fatalf("error looking up plugin symbol:%v", err)
	}

	ps, ok := symbol.(PluginSay)
	if !ok {
		log.Fatalf("symbol type not PluginSay:%T", symbol)
	}
	ps.Say("Plugin", "success")
}
