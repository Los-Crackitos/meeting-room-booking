#[derive(Serialize, Deserialize)]
pub struct Room {
    pub id: Option<i32>,
    pub name: String,
    pub capacity: i32
}
