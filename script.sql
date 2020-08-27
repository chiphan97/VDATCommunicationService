CREATE DATABASE dchat;

CREATE TABLE Groups(
    id_group serial ,
    sub_user_create varchar(100) NOT NULL,
    name_group varchar(100),
    type_group varchar(10) NOT NULL,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp,
    PRIMARY KEY (id_group)
);

CREATE TABLE Groups_Users(
    id_group serial ,
    sub_user_join varchar(100),
    last_deleted_messages timestamp not null default now(),
    join_at timestamp not null default now(),
    out_at timestamp,
    PRIMARY KEY (id_group,sub_user_join),
    CONSTRAINT FK_Groups_Groups_Users FOREIGN KEY (id_group) REFERENCES Groups (id_group)

);

CREATE TABLE Messages(
    id_mess serial,
    subject_sender varchar(100) ,
    content text,
    id_group serial NOT NULL ,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp,
    PRIMARY KEY (id_mess),
    CONSTRAINT FK_Groups_Messages FOREIGN KEY (id_group) REFERENCES Groups (id_group)
);

CREATE TABLE Messages_Delete(
    id_mess serial,
    sub_user_deleted varchar(100),
    created_at timestamp not null default now(),
    PRIMARY KEY (id_mess,sub_user_deleted),
    CONSTRAINT FK_Messages_Messages_Delete FOREIGN KEY (id_mess) REFERENCES Messages (id_mess)
);