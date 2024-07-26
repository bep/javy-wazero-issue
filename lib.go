package javywazeroissue

import (
	"bytes"
	"context"
	_ "embed"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

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

	compiled, err := r.CompileModule(ctx, fooWasm)
	if err != nil {
		return "", err
	}

	buff := &bytes.Buffer{}
	buff.WriteString(`{ "n": 2, "bar": "baz" }`)

	config := wazero.NewModuleConfig().WithStdout(buff).WithStderr(os.Stderr).WithStdin(buff)

	_, err = r.InstantiateModule(ctx, compiled, config)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
