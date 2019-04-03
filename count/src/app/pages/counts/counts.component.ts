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
  }

  sessionsMap: Map<String, Session[]> 
  sessions: Session[]
  error: any
  sessionIsClicked = false;

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

  goToCount(session: Session) {
    console.log(session)
    this.sessionIsClicked = true;
    
  }
}
