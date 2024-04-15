CREATE TABLE news (
    id serial primary KEY,
    article varchar(255) DEFAULT NULL,
    body text DEFAULT NULL
)


-- Create rating of news and commentary