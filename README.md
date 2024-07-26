


Build:

```
javy emit-provider -o quickjs.wasm
javy compile foo.js -d -o foo.wasm
```

Verify from CLI:

```bash
echo '{ "n": 2, "bar": "baz" }' | wasmtime --preload javy_quickjs_provider_v2=quickjs.wasm foo.wasm
```

Verify from Go:

```
go test
```