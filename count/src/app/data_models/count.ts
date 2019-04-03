export enum time {
  "beginning",
  "middle",
  "end"
}
export class Count {
  count: number;
  countID: number;
  time: time;
  userName: string;

  constructor(count: number, time: time, userName: string) {
    this.countID = -1;
    this.count = count;
    this.userName = userName;
    this.time = time;
  }
}
