extern crate diesel;

use crate::*;
use crate::models::*;

pub fn get_users() -> Vec<User> {
    use crate::schema::users::dsl::*;

    let connection = establish_database_connection();

    users.limit(5)
        .load::<User>(&connection)
        .expect("Error loading posts")
}
