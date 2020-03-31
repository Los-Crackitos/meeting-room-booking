#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use]
extern crate rocket;
extern crate rocket_contrib;
extern crate api;

use rocket_contrib::json::{Json};



use api::models::User;
use api::db::users::get_users;

#[get("/user")]
fn user() -> Json<Vec<User>> {
    Json(get_users())
}

fn main() {
    rocket::ignite()
        .mount("/api", routes![user])
        .launch();
}
