CREATE TABLE user (userid INT NOT NULL, username CHAR(30), xtoken VARCHAR(256) );

ALTER TABLE user ADD CONSTRAINT PK_user PRIMARY KEY (userid);

INSERT INTO user VALUES(1,"loman","123fjdsafjasdklkjkjkljljlkjkl");
INSERT INTO user VALUES(2,"ken","123fjdsafjasdklkjkjkljljlkjkl");
INSERT INTO user VALUES(3,"ken","ken");