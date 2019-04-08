import { Component, OnInit, NgModule } from "@angular/core";


import {Count} from '../../data_models/count';
import {CountService} from '../../services/count.service';
import {SpeakerService} from '../../services/speaker.service';
import {Speaker} from "../../data_models/speaker";
import {SessionService} from "../../services/session.service";
import { Session } from "../../data_models/session";


@Component({
  selector: 'app-counts',
  templateUrl: './counts.component.html',
  styleUrls: ['./counts.component.css']

})

export class CountsComponent implements OnInit {
  constructor(
    private countService: CountService,
    private speakerService: SpeakerService,
    private sessionService: SessionService
  ) {}

  speakers: Speaker[];
  sessions: Session[];
  error: any;
  buttonPressed: boolean = false;

ngOnInit() {
  this.getAllSpeakers();
  this.getAllSessions();
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

toggle(): void{
  if(this.buttonPressed){
    this.buttonPressed = false;
  } else {
    this.buttonPressed = true;
  }
}

}
