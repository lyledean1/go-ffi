use std::ffi::CStr;
use std::os::raw::c_char;

#[no_mangle]
pub extern "C" fn print_value(value: *const c_char) {
    let c_str = unsafe { CStr::from_ptr(value) };
    let str_slice = c_str.to_str().unwrap();
    println!("{}", str_slice);
}
