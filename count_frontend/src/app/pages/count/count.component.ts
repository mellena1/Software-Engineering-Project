import { Component, OnInit } from "@angular/core";

import { Count, timeMapping } from "../../data_models/count";
import { CountService } from "../../services/count.service";
import { SessionService } from "../../services/session.service";
import { Session } from "src/app/data_models/session";
import { LoginService } from "src/app/services/login.service";

@Component({
  selector: "app-counts",
  templateUrl: "./count.component.html",
  styleUrls: ["./count.component.css"]
})
export class CountComponent implements OnInit {
  constructor(
    private countService: CountService,
    private sessionService: SessionService,
    private loginService: LoginService
  ) {}

  ngOnInit() {
    this.getSessionsByTimeslot();
  }

  sessionsMap: Map<String, Session[]>;
  sessions: Session[];
  sessionIsClicked = false;
  selectedSession = new Session(null, null, null, null);
  model = new Count(null, null, this.loginService.getCurrentUsername());
  times = Object.values(timeMapping);
  error: any;

  onSubmit() {
    this.writeACount(this.model);
    this.sessionIsClicked = false;
  }

  getSessionsByTimeslot() {
    this.sessionService
      .getSessionsByTimeslot()
      .subscribe(
        data => (this.sessionsMap = data),
        error => (this.error = error)
      );
  }

  goToCount(session: Session) {
    this.sessionIsClicked = true;
    this.selectedSession = session;
    // resets model so fields are entry
    this.model = new Count(null, null, this.loginService.getCurrentUsername());
  }

  writeACount(count: Count) {
    this.countService
      .writeACount(this.selectedSession.id, count)
      .subscribe(
        response => console.log(response),
        error => (this.error = error)
      );
  }
}
