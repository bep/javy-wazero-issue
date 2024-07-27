package javywazeroissue

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed quickjs.wasm
var quickjsWasm []byte

//go:embed foo.wasm
var fooWasm []byte

func RunFoo() (string, error) {
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)

	// Instantiate WASI, which implements system I/O such as console output.
	if _, err := wasi_snapshot_preview1.Instantiate(ctx, r); err != nil {
		return "", err
	}

	compiledQuickJS, err := r.CompileModule(ctx, quickjsWasm)
	if err != nil {
		return "", err
	}

	compiledFoo, err := r.CompileModule(ctx, fooWasm)
	if err != nil {
		return "", err
	}

	stderr1 := &bytes.Buffer{}
	stderr2 := &bytes.Buffer{}

	stdinout1 := &bytes.Buffer{}
	stdinout2 := &bytes.Buffer{}

	stdinout1.WriteString(`{ "n": 2, "bar": "baz" }`)
	stdinout2.WriteString(`{ "n": 3, "bar": "foo" }`)

	_, err = r.InstantiateModule(ctx, compiledQuickJS, wazero.NewModuleConfig().
		WithStdout(stdinout1).
		WithStderr(stderr1).
		WithStdin(stdinout1).
		WithName("javy_quickjs_provider_v2"))
	if err != nil {
		return "", err
	}

	_, err = r.InstantiateModule(ctx, compiledFoo, wazero.NewModuleConfig().
		WithStdout(stdinout2).
		WithStderr(stderr2).
		WithStdin(stdinout2).
		WithName("foo"))
	if err != nil {
		return "", err
	}

	fmt.Println("stderr1:", stderr1.String())     // This prints nothing
	fmt.Println("stderr2:", stderr2.String())     // console.log from foo.js
	fmt.Println("stdinout1:", stdinout1.String()) // { "n": 2, "bar": "baz" }
	fmt.Println("stdinout2:", stdinout2.String()) // "foo":4,"newBar":"foo!"}

	return stdinout1.String(), nil
}
