import { Room } from "./room";
import { Speaker } from "./speaker";
import { Timeslot } from "./timeslot";

export class Session {
  id: number;
  name: string;
  room: Room;
  speaker: Speaker;
  timeslot: Timeslot;

  constructor(name: string, room: Room, speaker: Speaker, timeslot: Timeslot) {
    this.id = -1;
    this.name = name;
    this.room = room;
    this.speaker = speaker;
    this.timeslot = timeslot;
  }
}
