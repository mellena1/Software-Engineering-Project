USE codecamp;

INSERT INTO room (`roomName`, `capacity`) VALUES 
                 ('Gump', 21),
                 ('Jones', 0),
                 ('Wayne', 50),
                 ('Ripley', 17),
                 ('Max', 12);

INSERT INTO speaker (`email`, `firstName`, `lastName`) VALUES 
                    ('audrey.kirlin@example.org', 'Bernadette', 'Mante'),
                    ('conn.kelsi@example.net', 'Pat', 'Davis'),
                    ('dortha00@example.com', 'Adelia', 'Bogisich'),
                    ('haley.stevie@example.org', 'Yvonne', 'Gutmann'),
                    ('oconnell.obie@example.org', 'Viva', 'Pagac');

INSERT INTO session (`startTime`, `endTime`, `sessionName`, `email`, `roomName`) VALUES
                    ('2019-02-18T21:00:00', '2019-02-18T22:00:00', 'Microservices', 'audrey.kirlin@example.org', 'Gump'),
                    ('2019-02-18T14:00:00', '2019-02-18T15:00:00', 'Connected Devices', 'conn.kelsi@example.net', 'Jones'),
                    ('2019-02-18T11:00:00', '2019-02-18T12:30:00', 'Exploring Blockchain', 'dortha00@example.com', 'Wayne'),
                    ('2019-02-18T10:00:00', '2019-02-18T11:00:00', 'Clean Code Smean Code', 'haley.stevie@example.org', 'Ripley'),
                    ('2019-02-18T21:00:00', '2019-02-18T22:00:00', 'Bet You Didn\'t Think', 'oconnell.obie@example.org', 'Max');
