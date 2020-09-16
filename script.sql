CREATE DATABASE dchat;

CREATE TABLE Groups(
    id_group serial ,
    owner_id varchar(100) NOT NULL,
    name varchar(100),
    type varchar(15) NOT NULL,
    private boolean,
    thumbnail varchar(255),
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp,
    PRIMARY KEY (id_group)
);

CREATE TABLE Groups_Users(
    id_group serial,
    user_id varchar(100),
    join_at timestamp not null default now(),
    PRIMARY KEY (id_group,user_id),
    CONSTRAINT FK_Groups_Groups_Users FOREIGN KEY (id_group) REFERENCES Groups (id_group)
);

CREATE TABLE Messages(
    id_mess serial,
    user_sender varchar(100) ,
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
    user_deleted varchar(100),
    PRIMARY KEY (id_mess,user_deleted),
    CONSTRAINT FK_Messages_Messages_Delete FOREIGN KEY (id_mess) REFERENCES Messages (id_mess)
);

CREATE TABLE UserDetail(
    user_id  varchar(100),
    username varchar(100),
    first    varchar(100),
    last     varchar(100),
    role     varchar(15),
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp,
    PRIMARY KEY (user_id)
);
CREATE TABLE ONLINE(
    hostname varchar(100),
    socket_id varchar(100),
    user_id varchar(100) ,
    log_at timestamp not null default now(),
    CONSTRAINT FK_ONLINE_USER FOREIGN KEY (user_id) REFERENCES UserDetail (user_id),
    PRIMARY KEY (hostname,socket_id)
);
