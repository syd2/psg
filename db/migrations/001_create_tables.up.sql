
CREATE TABLE users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    username varchar(200) NOT NULL UNIQUE,
    password varchar(100) NOT NULL
);

CREATE TABLE passwords (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    app_name varchar(200) NOT NULL PRIMARY KEY,
    app_password varchar(100) NOT NULL,
    user_id BIGINT REFERENCES users (id) ON DELETE CASCADE NOT NULL
);