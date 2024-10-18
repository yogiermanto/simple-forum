ALTER TABLE users
ADD COLUMN username varchar(100) NOT NULL;

ALTER TABLE users
ADD CONSTRAINT UNIQUE unique_username (username);