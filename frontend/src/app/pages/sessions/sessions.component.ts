import { Component, OnInit, NgModule } from "@angular/core";
import { SessionService } from "../../services/session.service";
import { Session } from "../../data_models/session";
import { FormGroup, FormControl, Validators } from "@angular/forms";

@Component({
  selector: "app-sessions",
  templateUrl: "./sessions.component.html",
  styleUrls: ["./sessions.component.css"]
})
export class SessionsComponent implements OnInit {
  constructor(private sessionService: SessionService) {}
  sessions: Session[];
  session = new Session("", {id: 0, name: "", capacity: 0}, {id: 0, firstName: "", lastName: "", email: ""}, {id: 0, startTime: "", endTime: ""});
  selectedSession: Session;
  error: any;
  isEditable = false;
  sessionForm = new FormGroup({
    sessionName: new FormControl(""),
    sessionRoom: new FormControl(""),
    sessionSpeaker: new FormControl(""),
    sessionTimeslot: new FormControl("")
  });

  ngOnInit() {
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

  getSession(id: number) {
    this.sessionService
      .getSession(id)
  }

  updateSession(updatedSession: Session) {
    this.sessionService
      .updateSession(updatedSession)
  }

  deleteSession(id: number) {
    if (confirm("Are you sure you want to remove it?")) {
      this.sessionService
        .deleteSession(id)
        .subscribe(error => (this.error = error));
      console.log("The following Session Deleted :", this.sessionForm.value);
      this.sessions = this.sessions.filter(item => item.id !== id);
    }
  }

  onSelect(session: Session): void {
    this.selectedSession = session;
  }

  onSubmit(): void {
    var newSession = new Session(
      this.session.name,
      this.session.room,
      this.session.speaker,
      this.session.timeslot
    );
    this.sessionService
        .writeSession(
          this.session.name,
          this.session.room.id,
          this.session.speaker.id,
          this.session.timeslot.id
        )
        .subscribe(response => (newSession.id = response.id), error => (this.error = error));
      console.log("Session Submitted!", this.sessionForm.value);
      this.sessionForm.reset();
      this.sessions.push(newSession);
  }
}
