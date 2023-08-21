# Go ffi

Simple examples of using ffi in Go with Rust and Zig.

I've set up a more complex example of using Rust with Flutter via ffi, so this is more a mini PoC with Go. See [Flutter Whisper.cpp](https://github.com/lyledean1/flutter_whisper.cpp) and my talk on [ffi with Rust at Fluttercon](https://www.droidcon.com/2023/08/07/supercharging-your-flutter-apps-with-rust/)

## Go -> Rust via ffi 

Execute Rust code in Go via ffi. 

Located in `/go-rust-ffi`

## Go -> Zig via ffi 

Execute Zig code in Go via ffi. 

Located in `/go-zig-ffi`

## Run 

You will need Go, Rust and Zig installed to run all the examples.

Just run `make run` in each folder, this will build the Rust code or Zig Code to a shared library then execute FFI via Go to run the function.
