# go-ffi

Execute Rust code via FFI. 

I've set up a more complex example of using Rust with Flutter via ffi, so this is more a mini PoC. See [Flutter Whisper.cpp](https://github.com/lyledean1/flutter_whisper.cpp) and my talk on [ffi with Rust at Fluttercon](https://www.droidcon.com/2023/08/07/supercharging-your-flutter-apps-with-rust/)

## Run 

You will need Rust and Go installed (tested on Rust 1.71.1 and Go 1.21)

Just run `make run`, this will build the Rust code to a shared library then execute FFI via Go to run the function.
