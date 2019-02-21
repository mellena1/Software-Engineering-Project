DROP DATABASE IF EXISTS codecamp;
CREATE DATABASE codecamp;

USE codecamp;

DROP TABLE IF EXISTS session,
                     timeslot,
                     room,
                     speaker;

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

CREATE TABLE timeslot (
    timeslotID  INT     AUTO_INCREMENT NOT NULL,
    startTime   DATETIME,
    endTime     DATETIME,
    PRIMARY KEY (timeslotID)
);

CREATE TABLE session (
    sessionID       INT      AUTO_INCREMENT NOT NULL,
    speakerID       INT,
    roomID          INT,
    timeslotID      INT,
    sessionName     VARCHAR(32),
    email           VARCHAR(32),
    roomName        VARCHAR(32),
    FOREIGN KEY (speakerID)  REFERENCES speaker (speakerID),
    FOREIGN KEY (roomID)     REFERENCES room (roomID),
    FOREIGN KEY (timeslotID) REFERENCES timeslot (timeslotID),
    PRIMARY KEY (sessionID)
);
