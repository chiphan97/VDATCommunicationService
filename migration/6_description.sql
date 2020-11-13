Alter table Messages
alter column parentID integer;

Alter table Messages
alter column numChild integer default 0;

ALTER table Messages
add CONSTRAINT FK_Groups_Messages_Child FOREIGN KEY (parentID) REFERENCES Messages(id_mess);



