CREATE TABLE users
(
    id            serial primary key not null unique,
    username      varchar(255)       not null unique,
    password_hash varchar(255)       not null
);

CREATE TABLE folders
(
    id          serial primary key not null unique,
    folder_name varchar(255)       not null,
    FolderDate  date               not null
);

CREATE TABLE files
(
    id        serial primary key not null unique,
    file_name varchar(255)       not null,
    file_date date               not null,
    file_size varchar(255)       not null,
    file_path varchar(255)       not null
);

CREATE TABLE users_folders
(
    id        serial primary key not null unique,
    user_id   int                not null,
    folder_id int                not null,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (folder_id) REFERENCES folders (id) ON DELETE CASCADE
);

CREATE TABLE folders_files
(
    id        serial primary key not null unique,
    folder_id int                not null,
    file_id   int                not null,
    FOREIGN KEY (folder_id) REFERENCES folders (id) ON DELETE CASCADE,
    FOREIGN KEY (file_id) REFERENCES files (id) ON DELETE CASCADE
);