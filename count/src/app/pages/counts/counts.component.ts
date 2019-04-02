import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from '@angular/forms';

import {Count} from '../../data_models/count';
import {CountService} from '../../services/count.service';
import { SessionService } from '../../services/session.service';
import { Session } from "../../data_models/session";
import { THIS_EXPR } from '@angular/compiler/src/output/output_ast';

@Component({
  selector: 'app-counts',
  templateUrl: './counts.component.html'
})
export class CountsComponent implements OnInit {
  constructor(private countService: CountService, private sessionService: SessionService) {}

  ngOnInit() {
    console.log("Hello world!");
    this.getSessionsByTimeslot()
    console.log(this.sessionsMap)
  }

  sessionsMap: Map<String, Session>
  error: any

  getSessionsByTimeslot(): void {
    this.sessionService.
      getSessionsByTimeslot()
      .subscribe(
        data => (this.sessionsMap = data),
        error => (this.error = error)
      );
      //.subscribe(data => console.log(data));
  }
}
