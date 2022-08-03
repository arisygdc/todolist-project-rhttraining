CREATE TABLE auth {
    id          uuid NOT NULL PRIMARY KEY,
    username    VARCHAR(30) NOT NULL,
    password    VARCHAR(255) NOT NULL,
    email       VARCHAR(60) NOT NULL
};