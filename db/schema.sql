DROP DATABASE IF EXISTS codecamp;
CREATE DATABASE codecamp;

USE codecamp;

DROP TABLE IF EXISTS session,
                     speaker,
                     room;

CREATE TABLE speaker (
    firstName   VARCHAR(32),
    lastName    VARCHAR(32),
    email       VARCHAR(32),
    PRIMARY KEY (email)
);

CREATE TABLE room (
    roomName   VARCHAR(32)   NOT NULL,
    capacity   INT           NOT NULL,
    PRIMARY KEY (roomName)
);

CREATE TABLE session (
    startTime       TIMESTAMP     NOT NULL,
    endTime         TIMESTAMP     NOT NUll,
    sessionName     VARCHAR(32)   NOT NULL,
    email           VARCHAR(32)   NOT NULL,
    roomName        VARCHAR(32)   NOT NULL,
    KEY (email),
    KEY (roomName),
    FOREIGN KEY (email)       REFERENCES speaker (email) ON DELETE CASCADE,
    FOREIGN KEY (roomName)    REFERENCES room (roomName) ON DELETE CASCADE,
    PRIMARY KEY (roomName, startTime)
);
