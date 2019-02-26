import { Room } from "./room";
import { Speaker } from "./speaker";
import { Timeslot } from "./timeslot";

export class Session {
    id: number;
    name: string;
    room: Room;
    speaker: Speaker;
    timeslot: Timeslot;
}