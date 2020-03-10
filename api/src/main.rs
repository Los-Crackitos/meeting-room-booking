#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use] extern crate rocket;
#[macro_use] extern crate rocket_contrib;
#[macro_use] extern crate serde_derive;

use rocket_contrib::json::{Json, JsonValue};

mod room;
use room::{Room};


#[get("/")]
fn readAll() -> Json<JsonValue> {
    Json(json!([
        "hero 1", 
        "hero 2"
    ]))
}

#[get("/<id>")]
fn read(id: i32) -> Json<JsonValue> {
    Json(json!([
        "hero 1", 
        "hero 2"
    ]))
}

#[post("/", data = "<room>")]
fn createRoom(room: Json<Room>) -> Json<Room> {
    room
}


fn main() {
    rocket::ignite()
        .mount("/room", routes![readAll, createRoom])
        .mount("/rooms", routes![read])
        .launch();
}
