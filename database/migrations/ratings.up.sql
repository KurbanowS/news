CREATE TABLE ratings (
    id serial primary KEY
    news_id bigint not null references news on delete cascade
    likes bigint DEFAULT NULL
    dislikes bigint DEFAULT NULL
)