export enum time {
  "beginning",
  "middle",
  "end"
}
export const timeMapping = {
  [time.beginning]: "beginning",
  [time.middle]: "middle",
  [time.end]: "end"
}

export class Count {
  count: number;
  sessionID: number;
  time: time;
  userName: string;

  constructor(count: number, time: time, userName: string) {
    this.sessionID = -1;
    this.count = count;
    this.userName = userName;
    this.time = time;
  }
}
