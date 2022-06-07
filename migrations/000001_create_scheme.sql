-- +goose Up
create table leader_board (
                              id serial primary key,
                              chat_id integer not null,
                              user_id integer not null,
                              score integer default 0,
                              full_name varchar(200) default '',
                              user_name varchar(200) default '',
                              unique (chat_id, user_id)
);

create index chat_id_index on leader_board(chat_id);
create index user_id_index on leader_board(user_id);


create table question (
                          id serial primary key,
                          text text,
                          answer varchar(50)
);

-- +goose Down
drop table leader_board;
drop table question;