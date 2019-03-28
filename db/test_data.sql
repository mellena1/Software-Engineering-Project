USE codecamp;

INSERT INTO speaker (`email`, `firstName`, `lastName`) VALUES 
                    ('audrey.kirlin@example.org', 'Bernadette', 'Mante'),
                    ('conn.kelsi@example.net', 'Pat', 'Davis'),
                    ('dortha00@example.com', 'Adelia', 'Bogisich'),
                    ('haley.stevie@example.org', 'Yvonne', 'Gutmann'),
                    ('oconnell.obie@example.org', 'Viva', 'Pagac');

INSERT INTO room (`roomName`, `capacity`) VALUES 
                 ('Gump', 21),
                 ('Jones', 0),
                 ('Wayne', 50),
                 ('Ripley', 17),
                 ('Max', 12);

INSERT INTO timeslot (`startTime`, `endTime`) VALUES
                     ('2019-02-18T21:00:00', '2019-02-18T22:00:00'),
                     ('2019-02-18T14:00:00', '2019-02-18T15:00:00'),
                     ('2019-02-18T11:00:00', '2019-02-18T12:30:00'),
                     ('2019-02-18T10:00:00', '2019-02-18T11:00:00'),
                     ('2019-02-18T21:00:00', '2019-02-18T22:00:00');

INSERT INTO session (`speakerID`, `roomID`, `timeslotID`, `sessionName`) VALUES
                    (1, 5, 4, 'Clean Code Smean Code'),
                    (2, 1, 5, 'Microservices'),        
                    (3, 2, 1, 'Connected Devices'),
                    (4, 3, 2, 'Exploring Blockchain'),
                    (5, 4, 3, 'Bet You Didt Think');          
                            

INSERT INTO count (`time`, `sessionID`, `userName`, `count`) VALUES
                     ('beginning', 1, "Kenny", 5),
                     ('middle', 2, "Charles", 10),
                     ('end', 2, "Kenny", 15),
                     ('beginning', 3, "Kenny", 20),
                     ('middle', 3, "Andrew", 25),
                     ('end', 3, "Kenny", 30);
