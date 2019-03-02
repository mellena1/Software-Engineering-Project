import { Component, OnInit, NgModule } from '@angular/core';
import { SessionService } from '../../services/session.service'
import { Session} from '../../data_models/session'
import { Room } from 'src/app/data_models/room';
import { Speaker } from 'src/app/data_models/speaker';
import { Timeslot } from 'src/app/data_models/timeslot';
import { FormGroup, FormControl, Validators, ReactiveFormsModule } from '@angular/forms';
import { FormsModule }    from '@angular/forms';

@Component({
  selector: 'app-sessions',
  templateUrl: './sessions.component.html',
  styleUrls: [ './sessions.component.css' ]
})
@NgModule({
  imports: [
    FormsModule,
    ReactiveFormsModule
  ],
  declarations: [
    SessionsComponent,
  ]
})
export class SessionsComponent implements OnInit {
  constructor(private sessionService: SessionService) { }
  
  sessions: Session[];
  newSession: Session;
  selectedSession: Session;
  error: any;
  sessionForm: FormGroup;

  
  ngOnInit() {
    this.getAllSessions();
    this.sessionForm = new FormGroup({
      'sessionName': new FormControl(this.newSession.name, [
        Validators.required
      ]),
      'sessionRoom': new FormControl(this.newSession.room, [
        Validators.required
      ]),
      'sessionSpeaker': new FormControl(this.newSession.speaker, [
        Validators.required
      ]),
      'sessionTime': new FormControl(this.newSession.timeslot, [
        Validators.required
      ])
    });
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
