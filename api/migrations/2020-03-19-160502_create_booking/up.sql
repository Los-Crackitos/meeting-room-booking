CREATE TABLE booking
(
    id SERIAL PRIMARY KEY,
    start_date TIMESTAMPTZ,
    end_date TIMESTAMPTZ,
    room_id SERIAL,
    user_id SERIAL,

    CONSTRAINT fk_room FOREIGN KEY (room_id)
    REFERENCES rooms(id),

    CONSTRAINT fk_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
)
