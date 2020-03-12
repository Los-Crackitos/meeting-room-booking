#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use] extern crate rocket;
#[macro_use] extern crate rocket_contrib;

extern crate postgres;

use rocket_contrib::json::{Json, JsonValue};
use dotenv;
use std::format;
use postgres::{Connection, TlsMode};


#[get("/")]
fn read_all() -> Json<JsonValue> {
    Json(json!([
        "hero 1", 
        "hero 2"
    ]))
}

fn main() {
    
    dotenv::dotenv().ok(); // Load .env

    let db_host = dotenv::var("POSTGRES_DB_HOST").unwrap();
    let db_name = dotenv::var("POSTGRES_DB_NAME").unwrap();
    let db_user = dotenv::var("POSTGRES_DB_USER").unwrap();
    let db_password = dotenv::var("POSTGRES_DB_PASSWORD").unwrap();

    let conn = Connection::connect(format!("postgresql://{}:{}@{}/{}",
        db_user, db_password, db_host, db_name ), TlsMode::None)
        .expect("DB connection failed");

    conn.execute("CREATE TABLE IF NOT EXISTS rooms (
        id              SERIAL PRIMARY KEY,
        name            VARCHAR NOT NULL,
        capacity        NUMERIC
    )", &[]).unwrap();

    conn.execute("CREATE TABLE IF NOT EXISTS users (
        id              SERIAL PRIMARY KEY,
        name            VARCHAR NOT NULL
    )", &[]).unwrap();

    conn.execute("CREATE TABLE IF NOT EXISTS booking (
        id              SERIAL PRIMARY KEY,
        start_date      TIMESTAMPTZ,
        end_date        TIMESTAMPTZ,
        room_id         SERIAL,
        user_id         SERIAL,

        CONSTRAINT fk_room FOREIGN KEY (room_id)
            REFERENCES public.rooms (id),
       
        CONSTRAINT fk_user FOREIGN KEY (user_id)
            REFERENCES public.users (id)
    )", &[]).unwrap();


    rocket::ignite()
        .mount("/room", routes![read_all])
        .launch();
}
