export enum time {
  "beginning",
  "middle",
  "end"
}
export class Count {
  Count: number;
  SessionID: number;
  Time: time;
  UserName: string;

  constructor(count: number, time: time, userName: string) {
    this.SessionID = -1;
    this.Count = count;
    this.UserName = userName;
    this.Time = time;
  }
}
