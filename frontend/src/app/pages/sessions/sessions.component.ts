import { Component, OnInit } from '@angular/core';
import { SessionService } from '../../services/session.service'
import { Session} from '../../data_models/session'
import { Room } from 'src/app/data_models/room';
import { Speaker } from 'src/app/data_models/speaker';
import { Timeslot } from 'src/app/data_models/timeslot';

@Component({
  selector: 'app-sessions',
  templateUrl: './sessions.component.html',
  styleUrls: [ './sessions.component.css' ]
})
export class SessionsComponent implements OnInit {
  sessions: Session[];
  selectedSession: Session;
  error: any;
  constructor(private sessionService: SessionService ) { }


  public mock = [{name: 'one', age: 30 },{ name: 'two', age: 27 },{ name: 'three', age: 50 }];

  ngOnInit() {
    this.getAllSessions();
  }

  getAllSessions(): void {
    this.sessionService
      .getAllSessions()
      //.subscribe(
        //sessions => (this.sessions = sessions),
        //error => (this.error = error)
      //)
  }

  getSession(id: number): void {
    this.sessionService
      .getSession(id)
  }

  writeSession(): void {
    this.sessionService
      .writeSession()
  }

  updateSession(): void {
    this.sessionService
      .updateSession()
  }

  deleteSession(): void {
    this.sessionService
      .deleteSession()
  }

  onSelect(session: Session): void {
    this.selectedSession = session;
  }

}
