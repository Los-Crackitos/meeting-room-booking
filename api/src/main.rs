#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use] extern crate rocket;
#[macro_use] extern crate rocket_contrib;

use rocket_contrib::json::{Json, JsonValue};

#[get("/")]
fn read_all() -> Json<JsonValue> {
    Json(json!([
        "hero 1", 
        "hero 2"
    ]))
}

fn main() {
       
    rocket::ignite()
        .mount("/room", routes![read_all])
        .launch();
}
