CREATE TABLE users
(
    id         serial       not null unique,
    user_alias varchar(255) not null,
    chat_id    varchar(255) not null
);

CREATE TABLE subscription_entities
(
    id        serial not null unique,
    user_id   int    references users(id)   on delete cascade not null,
    source_id int    references sources(id) on delete cascade not null
);

CREATE TABLE sources
(
    id           serial       not null unique,
    name         varchar(255) not null,
    feed_url     varchar(255) not null unique
);

CREATE TABLE articles
(
    id           serial       not null unique,
    source_id    int          references sources(id) on delete cascade not null
    title        varchar(255) not null,
    link         varchar(255) not null,
    summary      text         not null,
    published_at timestamp    not null,
    created_at   timestamp    not null,
    posted       boolean      default false 
);