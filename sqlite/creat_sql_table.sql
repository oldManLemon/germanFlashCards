CREATE TABLE ger_dict (
	id INTERGER PRIMARY KEY,
   	ger_article TEXT NOT NULL,
	ger_word TEXT NOT NULL UNIQUE,
	eng_word TEXT NOT NULL,
	score FLOAT DEFAULT 50

);

CREATE TABLE ger_dict (id INTERGER PRIMARY KEY, ger_article TEXT NOT NULL, ger_word TEXT NOT NULL UNIQUE, eng_word TEXT NOT NULL, score INT);

-- ADD check data

INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('m', 'Vogel', 'bird');
INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('f', 'Tür', 'door');
INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('m', 'Tisch', 'table');
INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('f', 'Wurst', 'sausage');
INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('n', 'Bier', 'beer');
INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('m', 'Hund', 'dog');
INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('f', 'Katze', 'cat');
INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('m', 'Schrank', 'cupboard');
INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('n', 'Tempo', 'time');
INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('n', 'Handy', 'handheld two-way radio');
INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('f', 'Polizei', 'police');
INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('n', 'papier', 'paper');

select * from ger_dict;
select rowid, ger_word from ger_dict;


sqlite> .output ger_dict.db
sqlite> .dump ger_dict

-- When new DB needs to start 
.read ger_dict.db