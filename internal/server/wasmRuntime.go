package server

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

type WasmRuntime struct{}

func (wrtm *WasmRuntime) RunAdd(a, b uint64) uint64 {

	// Choose the context to use for function calls.
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	// Instantiate WASI, which implements host functions needed for TinyGo to
	// implement `panic`.
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	//create the wasm filepath
	//get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filepath := strings.Split(cwd, "faas/")[0] + "/wasm/add.wasm"

	//load the wasm file
	addWasm, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	// Instantiate the guest Wasm into the same runtime. It exports the `add`
	// function, implemented in WebAssembly.
	mod, err := r.Instantiate(ctx, addWasm)
	if err != nil {
		log.Panicf("failed to instantiate module: %v", err)
	}

	// Call the `add` function and print the results to the console.
	add := mod.ExportedFunction("add")
	results, err := add.Call(ctx, a, b)
	if err != nil {
		log.Panicf("failed to call add: %v", err)
	}

	fmt.Printf("%d + %d = %d\n", a, b, results[0])
	return results[0]
}
