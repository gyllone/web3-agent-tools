
extern crate proc_macro;

use proc_macro::TokenStream;
use quote::quote;
use syn::{parse_macro_input, DeriveInput, Data, Fields};
use ai_agent_tool::{CCompatible, Release};

#[proc_macro_derive(CCompatible)]
pub fn derive_c_compatible(input: TokenStream) -> TokenStream {
    let input = parse_macro_input!(input as DeriveInput);
    let name = input.ident;

    let expanded = quote! {
        impl CCompatible for #name {}
    };
    TokenStream::from(expanded)
}

#[proc_macro_derive(Release)]
pub fn derive_release(input: TokenStream) -> TokenStream {
    let input = parse_macro_input!(input as DeriveInput);
    let name = input.ident;

    let fields = match input.data {
        Data::Struct(data_struct) => {
            match data_struct.fields {
                Fields::Named(field) => field.named,
                _ => panic!("Release only supports structs with named fields."),
            }
        },
        _ => panic!("Release only supports structs."),
    };

    let do_release = fields.iter().map(|field| {
        let field_name = field.ident.as_ref().unwrap();
        quote! {
            self.#field_name.release();
        }
    });

    let expanded = quote! {
        impl Release for #name {
            fn release(&self) {
                #(#do_release)*
            }
        }
    };
    TokenStream::from(expanded)
}
