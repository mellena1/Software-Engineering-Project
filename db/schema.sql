DROP DATABASE IF EXISTS codecamp;
CREATE DATABASE codecamp;

USE codecamp;

DROP TABLE IF EXISTS session,
                     timeslot,
                     room,
                     speaker,
                     count,
                     user;

CREATE TABLE speaker (
    speakerID   INT          AUTO_INCREMENT NOT NULL,
    email       VARCHAR(32),
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
    startTime   DATETIME NOT NULL,
    endTime     DATETIME NOT NULL,
    PRIMARY KEY (timeslotID)
);

CREATE TABLE session (
    sessionID       INT      AUTO_INCREMENT NOT NULL,
    speakerID       INT,
    roomID          INT,
    timeslotID      INT,
    sessionName     VARCHAR(32),
    FOREIGN KEY (speakerID)  REFERENCES speaker (speakerID)     ON DELETE SET NULL,
    FOREIGN KEY (timeslotID) REFERENCES timeslot (timeslotID)   ON DELETE SET NULL,
    FOREIGN KEY (roomID)     REFERENCES room (roomID)           ON DELETE SET NULL,
    FOREIGN KEY (timeslotID) REFERENCES timeslot (timeslotID)   ON DELETE SET NULL,
    PRIMARY KEY (sessionID)
);


CREATE TABLE count (
    time      ENUM('beginning', 'middle', 'end'),
    count     INT,
    userName  VARCHAR(32),
    sessionID INT,
    FOREIGN KEY (sessionID) REFERENCES session (sessionID)  ON DELETE CASCADE,
    PRIMARY KEY (time, sessionID)
);


