USE shortener;

DROP TABLE IF EXISTS url_mapping;

CREATE TABLE `url_mapping` (
  `url`  VARCHAR(3000) NOT NULL,
  `hash` VARCHAR(200)  NOT NULL,
  PRIMARY KEY (`url`),
  KEY `search_index` (`url`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = latin1;

DROP TABLE IF EXISTS hash_mapping;

CREATE TABLE `hash_mapping` (
  `hash` VARCHAR(200)  NOT NULL,
  `url`  VARCHAR(3000) NOT NULL,
  PRIMARY KEY (`hash`),
  KEY `search_index` (`hash`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = latin1;