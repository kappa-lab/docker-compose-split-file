DROP SCHEMA IF EXISTS sample;
CREATE SCHEMA sample;
USE sample;

DROP TABLE IF EXISTS items;

CREATE TABLE items
(
  id           INT(10),
  name     VARCHAR(40)
);

INSERT INTO items (id, name) VALUES (1, "Coffee");
INSERT INTO items (id, name) VALUES (2, "Tea");