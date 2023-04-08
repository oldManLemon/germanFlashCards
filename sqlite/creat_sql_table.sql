CREATE TABLE ger_dict (
	id INTERGER PRIMARY KEY,
   	ger_article TEXT NOT NULL,
	ger_word TEXT NOT NULL UNIQUE,
	eng_word TEXT NOT NULL

);

CREATE TABLE ger_dict (id INTERGER PRIMARY KEY, ger_article TEXT NOT NULL, ger_word TEXT NOT NULL UNIQUE, eng_word TEXT NOT NULL);

-- ADD check data

INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('m', 'Tisch', 'table');
INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES ('f', 'Tür', 'door');

select * from ger_dict;
select rowid, ger_word from ger_dict;


sqlite> .output ger_dict.db
sqlite> .dump ger_dict

-- When new DB needs to start 
.read ger_dict.db