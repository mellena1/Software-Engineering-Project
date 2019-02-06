DROP DATABASE IF EXISTS codecamp;
CREATE DATABASE IF NOT EXISTS codecamp;

USE codecamp;

DROP TABLE IF EXISTS sessions,
                     speakers,
                     rooms;

CREATE TABLE 'speakers' (
    speakerID   INT           NOT NULL,
    name        VARCHAR(32)   NOT NULL,
    PRIMARY KEY (speakerID)
);

CREATE TABLE 'rooms' (
    roomID     INT   NOT NULL,
    capacity   INT   NOT NULL,
    PRIMARY KEY (roomID)
);

CREATE TABLE 'sessions' (
    startTime   TIMESTAMP     NOT NULL,
    endTIme     TIMESTAMP     NOT NUll,
    title       VARCHAR(32)   NOT NULL,
    speakerID   INT           NOT NULL,
    roomID      INT           NOT NULL,
    KEY (speakerID)
    KEY (roomID),
    FOREIGN KEY (speakerID) REFERENCES speakers (speakerID) ON DELETE CASCADE,
    FOREIGN KEY (roomID)    REFERENCES rooms    (roomID)    ON DELETE CASCADE,
    PRIMARY KEY (roomID, startTime)
);
