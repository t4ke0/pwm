-- DROP TABLE IF EXISTS sessions;
-- DROP TABLE IF EXISTS passwords;
-- DROP TABLE IF EXISTS user_t;
-- DROP TABLE IF EXISTS server_encryption_key;
-- DROP TABLE IF EXISTS server_auth_key;

CREATE TABLE IF NOT EXISTS user_t (
   id SERIAL PRIMARY KEY,
   username VARCHAR unique not null,
   password VARCHAR not null,
   email VARCHAR not null,
   key VARCHAR not null -- encryption and decryption user key.
);

CREATE TABLE IF NOT EXISTS sessions (
   session_id VARCHAR PRIMARY KEY,
   jwt_token VARCHAR not null,
   user_id INTEGER,
   created_at TIMESTAMP,
   revoked boolean,
   FOREIGN KEY(user_id) REFERENCES user_t(id)
);


CREATE TABLE IF NOT EXISTS passwords (
   id SERIAL PRIMARY KEY,
   user_id INTEGER not null,
   password VARCHAR not null,
   username VARCHAR,
   category TEXT,
   site TEXT,
   FOREIGN KEY(user_id) REFERENCES user_t(id)
);

CREATE TABLE IF NOT EXISTS server_encryption_key (
   server_key VARCHAR
);

CREATE TABLE IF NOT EXISTS server_auth_key (
   auth_server_key VARCHAR
);
