use serde_derive::{Deserialize, Serialize};

#[derive(Queryable, Deserialize, Serialize)]
pub struct User {
    pub id: i32,
    pub name: String,
}
