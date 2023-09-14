
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
           0: 0x1e6b - runtime.abort
                           at /tinygo/src/runtime/runtime_tinygowasm.go:70:6                 - runtime._panic
                           at /tinygo/src/runtime/panic.go:52:7
           1: 0x4645c - <unknown>!(*encoding/json.decodeState).literalStore
           2: 0x40f94 - (reflect.Value).Cap
                           at /usr/local/go/src/encoding/json/decode.go:388:28              - (*encoding/json.decodeState).readIndex
                           at /usr/local/go/src/encoding/json/decode.go:545:17              - (*encoding/json.decodeState).value
                           at /usr/local/go/src/encoding/json/decode.go:388:53
           3: 0x1a902 - (*encoding/json.decodeState).unmarshal
                           at /usr/local/go/src/encoding/json/decode.go:181:16              - encoding/json.Unmarshal
                           at /usr/local/go/src/encoding/json/decode.go:108:20
           4: 0x43506 - (*omb-wasm/avro.Node).UnmarshalJSON
                           at /home/rockwood/Workspace/tinygo-oob-memory-issue/avro/node.go:160:27              - (Go interface method)
                           at <Go interface method>
           5: 0x484a2 - <unknown>!(*encoding/json.decodeState).object
           6: 0x4102e - (reflect.Value).Cap
                           at /usr/local/go/src/encoding/json/decode.go:374:22              - (*encoding/json.decodeState).readIndex
                           at /usr/local/go/src/encoding/json/decode.go:545:17              - (*encoding/json.decodeState).value
                           at /usr/local/go/src/encoding/json/decode.go:388:53
           7: 0x40cf2 - (reflect.Value).Kind
                           at /usr/local/go/src/encoding/json/decode.go:555:21              - (*encoding/json.decodeState).array
                           at /usr/local/go/src/encoding/json/decode.go:544:12              - (*encoding/json.decodeState).value
                           at /usr/local/go/src/encoding/json/decode.go:364:21
           8: 0x1a902 - (*encoding/json.decodeState).unmarshal
                           at /usr/local/go/src/encoding/json/decode.go:181:16              - encoding/json.Unmarshal
                           at /usr/local/go/src/encoding/json/decode.go:108:20
           9: 0x435e0 - (*omb-wasm/avro.Node).UnmarshalJSON
                           at /home/rockwood/Workspace/tinygo-oob-memory-issue/avro/node.go:174:27              - (Go interface method)
                           at <Go interface method>
          10: 0x484a2 - <unknown>!(*encoding/json.decodeState).object
          11: 0x4102e - (reflect.Value).Cap
                           at /usr/local/go/src/encoding/json/decode.go:374:22              - (*encoding/json.decodeState).readIndex
                           at /usr/local/go/src/encoding/json/decode.go:545:17              - (*encoding/json.decodeState).value
                           at /usr/local/go/src/encoding/json/decode.go:388:53
          12: 0x40cf2 - (reflect.Value).Kind
                           at /usr/local/go/src/encoding/json/decode.go:555:21              - (*encoding/json.decodeState).array
                           at /usr/local/go/src/encoding/json/decode.go:544:12              - (*encoding/json.decodeState).value
                           at /usr/local/go/src/encoding/json/decode.go:364:21
          13: 0x1a902 - (*encoding/json.decodeState).unmarshal
                           at /usr/local/go/src/encoding/json/decode.go:181:16              - encoding/json.Unmarshal
                           at /usr/local/go/src/encoding/json/decode.go:108:20
          14: 0x435e0 - (*omb-wasm/avro.Node).UnmarshalJSON
                           at /home/rockwood/Workspace/tinygo-oob-memory-issue/avro/node.go:174:27              - (Go interface method)
                           at <Go interface method>
          15: 0x484a2 - <unknown>!(*encoding/json.decodeState).object
          16: 0x4102e - (reflect.Value).Cap
                           at /usr/local/go/src/encoding/json/decode.go:374:22              - (*encoding/json.decodeState).readIndex
                           at /usr/local/go/src/encoding/json/decode.go:545:17              - (*encoding/json.decodeState).value
                           at /usr/local/go/src/encoding/json/decode.go:388:53
          17: 0x1a902 - (*encoding/json.decodeState).unmarshal
                           at /usr/local/go/src/encoding/json/decode.go:181:16              - encoding/json.Unmarshal
                           at /usr/local/go/src/encoding/json/decode.go:108:20
          18: 0x19aef - (*omb-wasm/avro.Interop).UnmarshalJSON
                           at /home/rockwood/Workspace/tinygo-oob-memory-issue/avro/interop.go:531:27              - main.main
                           at /home/rockwood/Workspace/tinygo-oob-memory-issue/parse.go:14:24              - runtime.run$1
                           at /tinygo/src/runtime/scheduler_any.go:25:11
          19: 0x18752 - <goroutine wrapper>
                           at /tinygo/src/runtime/scheduler_any.go:23:2
          20:  0x49e - tinygo_launch
                           at /tinygo/src/internal/task/task_asyncify_wasm.S:59
          21: 0x185fa - runtime.run
                           at /tinygo/src/internal/task/task_asyncify.go:109:17              - _start
                           at /tinygo/src/runtime/runtime_wasm_wasi.go:21:5
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
