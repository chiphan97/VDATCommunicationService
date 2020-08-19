CREATE DATABASE dchat;

CREATE TABLE chatboxs(
    id serial,
    sender_id varchar(100) unique,
    receiver_id varchar(100) unique,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp,
    primary key (id)
);

CREATE TABLE messages(
    id_mess serial,
    id_chat serial,
    content text,
    seen_at timestamp,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp,
    CONSTRAINT fk_chatbox_message FOREIGN KEY (id_chat) REFERENCES chatboxs (id)
);
