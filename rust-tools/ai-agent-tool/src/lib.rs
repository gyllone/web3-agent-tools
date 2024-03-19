
use core::ffi::*;

pub trait CCompatible {}

macro_rules! impl_c_compatible {
    ($($t:ty),*) => {
        $(
            impl CCompatible for $t {}
            impl CCompatible for *const $t {}
            impl CCompatible for *mut $t {}
        )*
    }
}

impl_c_compatible!(
    c_char,
    c_uchar,
    c_short,
    c_ushort,
    c_int,
    c_uint,
    c_long,
    c_ulong,
    c_float,
    c_double,
    c_void
);

pub trait Release: CCompatible {
    fn release(&self) {}
}

macro_rules! impl_release {
    ($($t:ty),*) => {
        $(
            impl Release for $t {}
        )*
    }
}

impl_release!(
    c_char,
    c_uchar,
    c_short,
    c_ushort,
    c_int,
    c_uint,
    c_long,
    c_ulong,
    c_float,
    c_double,
    c_void
);

// #[proc_macro_derive(Release)]
// pub fn derive_release(input: TokenStream) -> TokenStream {
//     let input = parse_macro_input!(input as DeriveInput);
//     let name = input.ident;
//
//     let fields = match input.data {
//         Data::Struct(data_struct) => {
//             match data_struct.fields {
//                 Fields::Named(fields_named) => fields_named.named,
//                 _ => panic!("DoSomethingForAllFields only supports structs with named fields."),
//             }
//         },
//         _ => panic!("Release only supports structs."),
//     };
//     TokenStream::from(expanded)
// }

pub fn c_str_to_str<'a>(c_str: *const c_char) -> &'a str {
    unsafe { CStr::from_ptr(c_str).to_str().expect("not utf-8 str") }
}

#[repr(C)]
pub struct CInputList<T: CCompatible> {
    len: u32,
    values: *const T,
}

impl<T: CCompatible> CCompatible for CInputList<T> {}

impl<T: CCompatible> From<CInputList<T>> for &[T] {
    fn from(value: CInputList<T>) -> Self {
        let len = value.len as usize;
        unsafe { core::slice::from_raw_parts(value.values, len) }
    }
}

#[repr(C)]
pub struct COutputList<T: CCompatible> {
    len: u32,
    values: *mut T,
}

impl<T: CCompatible> CCompatible for COutputList<T> {}

impl<T: CCompatible> From<Vec<T>> for COutputList<T> {
    fn from(vec: Vec<T>) -> Self {
        let len = vec.len();
        let mut new_vec = Vec::with_capacity(len);
        new_vec.extend(vec);
        let values = new_vec.as_mut_ptr();
        core::mem::forget(new_vec);
        Self {
            len: len as u32,
            values,
        }
    }
}

impl<T: CCompatible> Release for COutputList<T> {
    fn release(&self) {
        unsafe {
            let len = self.len as usize;
            let _ = Vec::from_raw_parts(self.values, len, len);
        }
    }
}

#[repr(C)]
pub struct CInputDict<T: CCompatible> {
    len: u32,
    keys: *const *const c_char,
    values: *const T,
}

impl<T: CCompatible> CCompatible for CInputDict<T> {}

impl<T: CCompatible> From<CInputDict<T>> for Vec<(&str, &T)> {
    fn from(value: CInputDict<T>) -> Self {
        let len = value.len as usize;
        let keys = unsafe { core::slice::from_raw_parts(value.keys, len) };
        let values = unsafe { core::slice::from_raw_parts(value.values, len) };
        keys.iter()
            .zip(values.iter())
            .map(|(&key, value)| (c_str_to_str(key), value))
            .collect()
    }
}

#[repr(C)]
pub struct COutputDict<T: CCompatible> {
    len: u32,
    keys: *mut *mut c_char,
    values: *mut T,
}

impl<T: CCompatible> CCompatible for COutputDict<T> {}

impl<T: CCompatible> From<Vec<(String, T)>> for COutputDict<T> {
    fn from(vec: Vec<(String, T)>) -> Self {
        let len = vec.len();
        let (keys, values): (Vec<_>, Vec<_>) = vec.into_iter().unzip();

        let mut new_keys = Vec::with_capacity(len);
        new_keys.extend(
            keys.into_iter().map(|key| {
                let mut key = key.into_bytes();
                key.reserve_exact(1);
                key.push(0);
                let key_ptr = key.as_mut_ptr() as *mut c_char;
                core::mem::forget(key);
                key_ptr
            })
        );
        let keys = new_keys.as_mut_ptr();
        core::mem::forget(new_keys);

        let mut new_values = Vec::with_capacity(len);
        new_values.extend(values);
        let values = new_values.as_mut_ptr();
        core::mem::forget(new_values);

        Self {
            len: len as u32,
            keys,
            values,
        }
    }
}

impl<T: CCompatible> Release for COutputDict<T> {
    fn release(&self) {
        unsafe {
            let len = self.len as usize;
            let keys = Vec::from_raw_parts(self.keys, len, len);
            for key in keys {
                let len = strlen(key) + 1; // Including the NUL byte
                let slice = core::slice::from_raw_parts_mut(key, len);
                let _ = Box::from_raw(slice as *mut [c_char] as *mut [u8]);
            }
            let _ = Vec::from_raw_parts(self.values, len, len);
        }
    }
}

extern "C" {
    /// Provided by libc or compiler_builtins.
    fn strlen(s: *const c_char) -> usize;
}