USE shortener;

DROP TABLE IF EXISTS url_mapping;

CREATE TABLE `url_mapping` (
  `hash`  VARCHAR(200)  NOT NULL,
  `value` VARCHAR(3000) NOT NULL,
  PRIMARY KEY (`hash`),
  KEY `search_index` (`hash`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = latin1;