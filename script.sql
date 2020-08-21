CREATE DATABASE dchat;

CREATE TABLE chatboxs(
    id serial,
    from_user varchar(100) unique,
    to_user varchar(100) unique,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp,
    primary key (id)
);

CREATE TABLE messages(
    id_mess serial,
    id_chat serial,
    body text,
    seen_at timestamp,
    sender_id varchar(100),
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp,
    CONSTRAINT fk_chatbox_message FOREIGN KEY (id_chat) REFERENCES chatboxs (id)
);
