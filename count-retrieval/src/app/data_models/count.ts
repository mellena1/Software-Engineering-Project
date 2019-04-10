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
