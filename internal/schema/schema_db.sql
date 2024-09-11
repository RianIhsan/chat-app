CREATE TABLE mst_user (
                          id serial primary key ,
                          username varchar(50) unique not null,
                          email varchar(100) unique not null,
                          password varchar(255) not null,
                          created_at timestamptz default now(),
                          updated_at timestamptz default now()
);

CREATE TABLE mst_friends (
                             id serial primary key,
                             user_id int references mst_user(id) on delete cascade ,
                             friend_id int references mst_user(id) on delete cascade,
                             status varchar(50) not null default 'pending',
                             created_at timestamptz default now()
);

CREATE TABLE mst_chats (
                           id serial primary key,
                           sender_id int references mst_user(id) on delete cascade,
                           receiver_id int references mst_user(id) on delete cascade,
                           message text not null ,
                           created_at timestamptz default now()
);

CREATE TABLE mst_chat_rooms (
                                id serial primary key,
                                name varchar(50) not null,
                                created_at timestamptz default now()
);

CREATE TABLE mst_room_members (
                                  room_id int references mst_chat_rooms(id) on delete cascade,
                                  user_id int references mst_user(id) on delete cascade,
                                  joined_at timestamptz default now(),
                                  primary key (room_id, user_id)
);

CREATE TABLE mst_tokens (
                            id serial primary key,
                            user_id int references mst_user(id) on delete cascade,
                            token varchar(255) not null,
                            expired_at timestamptz not null,
                            created_at timestamptz default now()
);