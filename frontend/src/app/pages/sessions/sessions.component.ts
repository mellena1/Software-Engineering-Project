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
  constructor(private sessionService: SessionService) { }
  sessions: Session[];
  selectedSession: Session;
  error: any;
  
  ngOnInit() {
    this.getAllSessions();
  }

  getAllSessions(): void {
    this.sessionService
      .getAllSessions()
      .subscribe(
        sessions => (this.sessions = sessions),
        error => (this.error = error)
      )
  }

  onSelect(session: Session): void {
    this.selectedSession = session;
  }
}
