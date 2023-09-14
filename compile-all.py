#!/bin/env python3

import subprocess

opt_levels = [
    '0',
    '1',
    '2',
    's',
    'z',
]

schedulers = ['none', 'asyncify']

for opt in opt_levels:
    for scheduler in schedulers:
        subprocess.check_call([
            "tinygo",
            "build",
            "-target=wasi",
            f"-opt={opt}",
            "-panic=print",
            f"-scheduler={scheduler}",
            f"-o=parse-{opt}-{scheduler}.wasm",
            "parse.go",
        ])

for opt in opt_levels:
    for scheduler in schedulers:
        returncode = subprocess.call([
            "wasmtime",
            f"parse-{opt}-{scheduler}.wasm",
        ])
        print(f"optimization level: {opt} scheduler: {scheduler} success: {returncode == 0}")
