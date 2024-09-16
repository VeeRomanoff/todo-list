CREATE TABLE users
(
    "uuid"          uuid         NOT NULL DEFAULT () PRIMARY KEY,
    "email"         varchar(255) NOT NULL DEFAULT '' UNIQUE,
    "username"      varchar(16)  NOT NULL UNIQUE,
    "password_hash" varchar(255) NOT NULL
);

CREATE TABLE todo_lists
(
    "uuid"  uuid         NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    "title" varchar(255) not null,
    "desc"  varchar(255) not null
);

CREATE TABLE users_lists
(
    "uuid"    uuid                                                NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    "user_id" uuid REFERENCES users (uuid) ON DELETE CASCADE      NOT NULL,
    "list_id" uuid REFERENCES todo_lists (uuid) ON DELETE CASCADE NOT NULL,
);

CREATE TABLE todo_items
(
    "uuid"  uuid         NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    "title" varchar(255) not null,
    "desc"  varchar(255) not null,
    "done"  boolean      not null default false
);

CREATE TABLE lists_items
(
    "uuid"      uuid                                                NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    "item_uuid" uuid REFERENCES todo_items (uuid) ON DELETE CASCADE NOT NULL,
    "list_id" uuid REFERENCES todo_lists (uuid) ON DELETE CASCADE NOT NULL,
);

