DROP DATABASE IF EXISTS codecamp;
CREATE DATABASE codecamp;

USE codecamp;

DROP TABLE IF EXISTS session,
                     speaker,
                     room;

CREATE TABLE speaker (
    speakerID   INT          AUTO_INCREMENT NOT NULL,
    email       VARCHAR(32)  NOT NULL,
    firstName   VARCHAR(32),
    lastName    VARCHAR(32),
    PRIMARY KEY (speakerID)
);

CREATE TABLE room (
    roomID     INT           AUTO_INCREMENT NOT NULL,
    roomName   VARCHAR(32)   NOT NULL,
    capacity   INT,
    PRIMARY KEY (roomID)
);

CREATE TABLE session (
    sessionID       INT      AUTO_INCREMENT NOT NULL,
    speakerID       INT,
    roomID          INT,
    startTime       DATETIME,
    endTime         DATETIME,
    sessionName     VARCHAR(32),
    email           VARCHAR(32),
    roomName        VARCHAR(32),
    FOREIGN KEY (speakerID) REFERENCES speaker (speakerID),
    FOREIGN KEY (roomID)    REFERENCES room (roomID),
    PRIMARY KEY (sessionID)
);
