import { Component, OnInit } from '@angular/core';
import { SessionService } from '../../services/session.service'
import { Session} from '../../data_models/session'

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


  ngOnInit() {
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
