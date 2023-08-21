build_rust:
	cd ffi_rs && cargo build --release

run: build_rust 
run:
	go run main.go