import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from '@angular/forms';

import { Count } from '../../data_models/count';
import { CountService } from '../../services/count.service';
import { SessionService } from '../../services/session.service'
import { Session } from 'src/app/data_models/session';

@Component({
  selector: 'app-counts',
  templateUrl: './counts.component.html'
})
export class CountsComponent implements OnInit {
  constructor(private countService: CountService, private sessionService: SessionService) {}

  ngOnInit() {
    this.getSessionsByTimeslot()
    this.getAllSessions()
  }

  sessionsMap: Map<String, Session[]>;
  sessions: Session[]
  error: any

  getSessionsByTimeslot() {
    this.sessionService.getSessionsByTimeslot()
      .subscribe(
        data => {
          console.log(data)
          this.sessionsMap = data
        },
        error => (this.error = error)
      );
  }

  getMapEntries() {
    return Array.from(this.sessionsMap.entries())
  }

  getAllSessions() {
    return this.sessionService.getAllSessions().subscribe(data => this.sessions = data)
  }
}
