CREATE TABLE users
(
    id         UUID        NOT NULL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name  VARCHAR(40) NOT NULL,
    auth_id    UUID        NOT NULL
);