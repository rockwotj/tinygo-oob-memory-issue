
Run `./compile-all.py` to compile using tinygo all the variations and then run all the compiled files.

On my machine the output is as follows:

```
wasm able to read json
optimization level: 0 scheduler: none success: True
wasm able to read json
optimization level: 0 scheduler: asyncify success: True
wasm able to read json
optimization level: 1 scheduler: none success: True
wasm able to read json
optimization level: 1 scheduler: asyncify success: True
wasm able to read json
optimization level: 2 scheduler: none success: True
panic: JSON decoder out of sync - data changing underfoot?
Error: failed to run main module `parse-2-asyncify.wasm`

Caused by:
    0: failed to invoke command default
    1: error while executing at wasm backtrace:
           0: 0x1e6b - <unknown>!runtime._panic
           1: 0x4645c - <unknown>!(*encoding/json.decodeState).literalStore
           2: 0x40f94 - <unknown>!(*encoding/json.decodeState).value
           3: 0x1a902 - <unknown>!encoding/json.Unmarshal
           4: 0x43506 - <unknown>!interface:{UnmarshalJSON:func:{slice:basic:uint8}{named:error}}.UnmarshalJSON$invoke
           5: 0x484a2 - <unknown>!(*encoding/json.decodeState).object
           6: 0x4102e - <unknown>!(*encoding/json.decodeState).value
           7: 0x40cf2 - <unknown>!(*encoding/json.decodeState).value
           8: 0x1a902 - <unknown>!encoding/json.Unmarshal
           9: 0x435e0 - <unknown>!interface:{UnmarshalJSON:func:{slice:basic:uint8}{named:error}}.UnmarshalJSON$invoke
          10: 0x484a2 - <unknown>!(*encoding/json.decodeState).object
          11: 0x4102e - <unknown>!(*encoding/json.decodeState).value
          12: 0x40cf2 - <unknown>!(*encoding/json.decodeState).value
          13: 0x1a902 - <unknown>!encoding/json.Unmarshal
          14: 0x435e0 - <unknown>!interface:{UnmarshalJSON:func:{slice:basic:uint8}{named:error}}.UnmarshalJSON$invoke
          15: 0x484a2 - <unknown>!(*encoding/json.decodeState).object
          16: 0x4102e - <unknown>!(*encoding/json.decodeState).value
          17: 0x1a902 - <unknown>!encoding/json.Unmarshal
          18: 0x19aef - <unknown>!runtime.run$1
          19: 0x18752 - <unknown>!runtime.run$1$gowrapper
          20:  0x49e - <unknown>!tinygo_launch
          21: 0x185fa - <unknown>!_start
       note: using the `WASMTIME_BACKTRACE_DETAILS=1` environment variable may show more debugging information
    2: wasm trap: wasm `unreachable` instruction executed
optimization level: 2 scheduler: asyncify success: False
wasm able to read json
optimization level: s scheduler: none success: True
wasm able to read json
optimization level: s scheduler: asyncify success: True
wasm able to read json
optimization level: z scheduler: none success: True
wasm able to read json
optimization level: z scheduler: asyncify success: True
```
