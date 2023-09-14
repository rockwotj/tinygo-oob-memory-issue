package main

import (
	_ "embed"

	"omb-wasm/avro"
)

//go:embed payload-1Kb.json
var payload []byte

func main() {
	i := avro.NewInterop()
	err := i.UnmarshalJSON(payload)
	if err != nil {
		println("error reading json:", err)
	} else {
		println("wasm able to read json")
	}
}
