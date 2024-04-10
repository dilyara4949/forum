create table users (
    email varchar primary key,
    username varchar,
    password text
);

create table categories (
    name varchar primary key
);

create table post (
    id serial primary key,
    email varchar REFERENCES users(email) on delete CASCADE,
    title text,
    body text,
    category varchar REFERENCES categories(name) on delete CASCADE
);

create table like_post_user  (
    email varchar REFERENCES users(email) on delete CASCADE,
    post_id int REFERENCES post(id) on delete CASCADE,
    primary key (email, post_id)
);

create table comments (
    email varchar REFERENCES users(email) on delete CASCADE,
    post_id int REFERENCES post(id) on delete CASCADE,
    message text
);
