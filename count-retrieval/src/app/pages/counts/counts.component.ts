import { Component, OnInit, NgModule } from "@angular/core";


import {Count, /*SpeakerCount, SessionCount*/} from '../../data_models/count';
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
  //sesssionsBySpeakers: SpeakerCount[];
  //countsBySessions: SessionCount[];
  counts: Count[];
  speakerSessionMap: Map<String, Map<String, Count[]>>;
  error: any;
  buttonPressed: boolean = false;
  speakerSelected: boolean = false;
  sessionSelected: boolean = false;

  //sessionBySpeaker = new SpeakerCount("",0,0,0);
  //countBySession = new SessionCount("",0,0,0);

  

  
ngOnInit() {
  this.getAllSpeakers();
  this.getAllSessions();
  //this.getAllCounts();
  this.getSessionsBySpeaker();
  //console.log("Here is sessionsBySpeakers: " + this.sesssionsBySpeakers);
  //console.log("Here is speakerSessionMap: " + this.speakerSessionMap);
  //var test = this.speakerSessionMap.keys();
  //console.log("Here are the keys: " + test);
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

/*getAllCounts(): void {
  this.countService
      .getAllCounts()
      .subscribe(
        counts => (this.counts = counts),
        error => (this.error = error)
      );
}*/

getSessionsBySpeaker(): void{
  this.countService
    .getCountsBySpeaker()
    .subscribe(
      speakerSessionMap => (this.speakerSessionMap = speakerSessionMap),
      error => (this.error = error)
    );
}

createTable(): void{
  if(this.buttonPressed){
    this.buttonPressed = false;
  } else {
    this.buttonPressed = true;
  }
}

submitSession(): void {
  this.sessionSelected = true;
  this.speakerSelected = false;
  //var newCountBySession = new SessionCount (this.speakerSessionMap.getKey(), 0,0,0);
  //console.log("Check it: " + newCountBySession);
}

submitSpeaker(): void {
  this.speakerSelected = true;
  this.sessionSelected = false;
  //console.log("Here is speakerSessionMap: " + this.speakerSessionMap);
  //console.log("Here is sessionsBySpeakers: " + this.sesssionsBySpeakers);
}

}
