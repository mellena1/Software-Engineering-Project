import { Component, OnInit } from "@angular/core";

import { Count, time } from "../../data_models/count";
import { CountService } from "../../services/count.service";
import { SpeakerService } from "../../services/speaker.service";
import { Speaker } from "../../data_models/speaker";
import { SessionService } from "../../services/session.service";
import { Session } from "../../data_models/session";

@Component({
  selector: "app-counts",
  templateUrl: "./counts.component.html",
  styleUrls: ["./counts.component.css"]
})
export class CountsComponent implements OnInit {
  constructor(
    private countService: CountService,
    private speakerService: SpeakerService,
    private sessionService: SessionService
  ) {}

  speakers: Speaker[];
  sessions: Session[];
  speakerSessionMap: Map<String, Map<String, Count[]>>;

  selectedSpeaker: string;
  selectedSession: Session;
  selectedSessionCounts: Count[];

  error: any;
  showTable: boolean = false;
  speakerSelected: boolean = false;
  sessionSelected: boolean = false;

  ngOnInit() {
    this.getAllSpeakers();
    this.getAllSessions();
    this.getSessionsBySpeaker();
  }

  getAllSessions(): void {
    this.sessionService
      .getAllSessions()
      .subscribe(
        sessions => (this.sessions = sessions),
        error => (this.error = error)
      );
  }

  getAllSpeakers(): void {
    this.speakerService
      .getAllSpeakers()
      .subscribe(
        speakers => (this.speakers = speakers),
        error => (this.error = error)
      );
  }

  getSessionsBySpeaker(): void {
    this.countService.getCountsBySpeaker().subscribe(
      speakerSessionMap => {
        this.speakerSessionMap = speakerSessionMap;
      },
      error => (this.error = error)
    );
  }

  getCountBySessionID(sessionID: number) {
    this.countService.getACount(sessionID).subscribe(
      (counts: Count[]) => {
        this.selectedSessionCounts = counts;
      },
      error => {
        this.error = error;
      }
    );
  }

  findCountFromList(countList: Array<Count>, timeWanted: string): string {
    for (var i = 0; i < countList.length; i++) {
      var count = countList[i];
      if (time[count.Time] === time[timeWanted]) {
        return `${count.Count.toString()} (${count.UserName})`;
      }
    }
    return "-";
  }

  submitSession(): void {
    this.showTable = true;
    this.sessionSelected = true;
    this.speakerSelected = false;
  }

  submitSpeaker(): void {
    this.showTable = true;
    this.speakerSelected = true;
    this.sessionSelected = false;
  }
}
