export enum time {
  "beginning",
  "middle",
  "end"
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

/*export class SessionCount {
  sessionName: string;
  countBeginning: number;
  countMiddle: number;
  countEnd: number;

  constructor(sessionName: string, countBeginning: number, countMiddle: number, countEnd: number) {
    this.sessionName = sessionName;
    this.countBeginning = countBeginning;
    this.countMiddle = countMiddle;
    this.countEnd = countEnd;
  }
}

export class SpeakerCount {
  speakerName: string;
  countBeginning: number;
  countMiddle: number;
  countEnd: number;

  constructor(speakerName: string, countBeginning: number, countMiddle: number, countEnd: number) {
    this.speakerName = speakerName;
    this.countBeginning = countBeginning;
    this.countMiddle = countMiddle;
    this.countEnd = countEnd;
  }
}*/
