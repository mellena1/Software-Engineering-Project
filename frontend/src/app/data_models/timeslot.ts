export class Timeslot {
  id: number;
  startTime: string;
  endTime: string;

  constructor(startTime: string, endTime: string) {
    this.id = -1;
    this.startTime = startTime;
    this.endTime = endTime;
  }
}
