CREATE SCHEMA urlshortener;
USE urlshortener;

DROP TABLE IF EXISTS urls;
CREATE TABLE urls (
  _id         INT AUTO_INCREMENT NOT NULL,
  urlid       VARCHAR(32) UNIQUE NOT NULL,
  longurl     TEXT NOT NULL,
  PRIMARY KEY (`_id`),
  INDEX `urlid_index` (`urlid`)
);

INSERT INTO urls 
  (urlid, longurl) 
VALUES 
  ('google', 'https://google.com'),
  ('shreyas', 'https://shreyasgupta.in'),
  ('youtube', 'https://www.youtube.com/watch?v=HpHdXfGSrN4');