DROP DATABASE IF EXISTS codecamp;
CREATE DATABASE IF NOT EXISTS codecamp;

USE codecamp;

DROP TABLE IF EXISTS session,
                     speaker,
                     room;

CREATE TABLE speaker (
    speakerID   INT,
    name        VARCHAR(32),
    PRIMARY KEY (speakerID)
);

CREATE TABLE room (
    roomID     INT   NOT NULL,
    capacity   INT   NOT NULL,
    PRIMARY KEY (roomID)
);

CREATE TABLE session (
    startTime       TIMESTAMP     NOT NULL,
    endTime         TIMESTAMP     NOT NUll,
    sessionName     VARCHAR(32)   NOT NULL,
    speakerID       INT           NOT NULL,
    roomID          INT           NOT NULL,
    KEY (speakerID),
    KEY (roomID),
    FOREIGN KEY (speakerID) REFERENCES speaker (speakerID) ON DELETE CASCADE,
    FOREIGN KEY (roomID)    REFERENCES room    (roomID)    ON DELETE CASCADE,
    PRIMARY KEY (roomID, startTime)
);
