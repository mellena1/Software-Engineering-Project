import { InMemoryDbService } from 'angular-in-memory-web-api';
import { Speaker } from '../data_models/speaker'
import { Room } from '../data_models/room';
import { Timeslot } from '../data_models/timeslot';
import { Session } from '../data_models/session';

export class MockApi implements InMemoryDbService {
    createDb() {
        const speakers: Speaker[] = [
            { id: 1, email: 'audrey.kirlin@example.org', firstName: 'Bernadette', lastName: 'Mante' },
            { id: 2, email: 'conn.kelsi@example.net',    firstName: 'Pat',        lastName: 'Davis' },
            { id: 3, email: 'dortha00@example.com',      firstName: 'Adelia',     lastName: 'Bogisich' },
            { id: 4, email: 'haley.stevie@example.org',  firstName: 'Yvonne',     lastName: 'Gutmann' },
            { id: 5, email: 'oconnell.obie@example.org', firstName: 'Viva',       lastName: 'Pagac' }
        ];

        const rooms: Room[] = [
            { id: 1, name: 'Gump',   capacity: 21 },
            { id: 2, name: 'Jones',  capacity: 0 },
            { id: 3, name: 'Wayne',  capacity: 50 },
            { id: 4, name: 'Ripley', capacity: 17 },
            { id: 5, name: 'Max',    capacity: 12 }
        ];
        
        const timeslots: Timeslot[] = [
            { id: 1, startTime: '2019-02-18T21:00:00', endTime: '2019-02-18T22:00:00' },
            { id: 2, startTime: '2019-02-18T14:00:00', endTime: '2019-02-18T15:00:00' },
            { id: 3, startTime: '2019-02-18T11:00:00', endTime: '2019-02-18T12:30:00' },
            { id: 4, startTime: '2019-02-18T10:00:00', endTime: '2019-02-18T11:00:00' },
            { id: 5, startTime: '2019-02-18T21:00:00', endTime: '2019-02-18T22:00:00' }
        ];

        const sessions: Session[] = [
            { id: 1, room: rooms[0], speaker: speakers[4], timeslot: timeslots[3], name: 'Clean Code Smean Code' } ,
            { id: 2, room: rooms[1], speaker: speakers[0], timeslot: timeslots[4], name: 'Microservices' } ,
            { id: 3, room: rooms[2], speaker: speakers[1], timeslot: timeslots[0], name: 'Connected Devices' } ,
            { id: 4, room: rooms[3], speaker: speakers[2], timeslot: timeslots[1], name: 'Exploring Blockchain' } ,
            { id: 5, room: rooms[4], speaker: speakers[3], timeslot: timeslots[2], name: 'Bet You Didn\'t Think' }
        ];
        
    return{rooms, speakers, timeslots, sessions}
    }
}