table! {
  booking (id) {
    id -> Int4,
    start_date -> Nullable<Timestamptz>,
    end_date -> Nullable<Timestamptz>,
    room_id -> Int4,
    user_id -> Int4,
  }
}

table! {
  rooms (id) {
    id -> Int4,
    name -> Varchar,
    capacity -> Nullable<Numeric>,
  }
}

table! {
  users (id) {
    id -> Int4,
    name -> Varchar,
  }
}

joinable!(booking -> rooms (room_id));
joinable!(booking -> users (user_id));

allow_tables_to_appear_in_same_query!(
  booking,
  rooms,
  users,
);
