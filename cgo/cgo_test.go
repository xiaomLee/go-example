package main

import "testing"

func TestSay(t *testing.T) {
	Say("Tom", "what a perfect day.")
}

func TestBye(t *testing.T) {
	Bye("Tom", "Jack")
}
