DROP DATABASE IF EXISTS codecamp;
CREATE DATABASE codecamp;

USE codecamp;

DROP TABLE IF EXISTS session,
                     speaker,
                     room;

CREATE TABLE speaker (
    email       VARCHAR(32)  NOT NULL,
    firstName   VARCHAR(32),
    lastName    VARCHAR(32),
    PRIMARY KEY (email)
);

CREATE TABLE room (
    roomName   VARCHAR(32)   NOT NULL,
    capacity   INT,
    PRIMARY KEY (roomName)
);

CREATE TABLE session (
    sessionID       INT      NOT NULL,
    startTime       TIMESTAMP,
    endTime         TIMESTAMP,
    sessionName     VARCHAR(32),
    email           VARCHAR(32),
    roomName        VARCHAR(32),
    KEY (email),
    KEY (roomName),
    FOREIGN KEY (email)       REFERENCES speaker (email),
    FOREIGN KEY (roomName)    REFERENCES room (roomName),
    PRIMARY KEY (sessionID)
);
