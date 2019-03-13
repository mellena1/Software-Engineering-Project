import { Component, OnInit, NgModule } from "@angular/core";
import { SessionService } from "../../services/session.service";
import { Session } from "../../data_models/session";
import { Room } from "src/app/data_models/room";
import { Speaker } from "src/app/data_models/speaker";
import { Timeslot } from "src/app/data_models/timeslot";
import { RoomService } from "src/app/services/room.service";
import { SpeakerService } from "src/app/services/speaker.service";
import { TimeslotService } from "src/app/services/timeslot.service";
import {
  FormGroup,
  FormControl,
  Validators,
  ReactiveFormsModule
} from "@angular/forms";
import { FormsModule } from "@angular/forms";

@Component({
  selector: "app-sessions",
  templateUrl: "./sessions.component.html",
  styleUrls: ["./sessions.component.css"]
})
@NgModule({
  imports: [FormsModule, ReactiveFormsModule],
  declarations: [SessionsComponent]
})
export class SessionsComponent implements OnInit {
  constructor(private sessionService: SessionService) {}
  
  sessions: Session[];
  session = new Session("", null, null, null);
  selectedSession: Session;
  error: any;
  sessionForm: FormGroup;

  ngOnInit() {
    this.getAllSessions();
    this.sessionForm = new FormGroup({
      name: new FormControl(""),
      room: new FormControl(""),
      speaker: new FormControl(""),
      timeslot: new FormControl("")
    });
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
        .subscribe(error => (this.error = error), id => (newSession.id = id));
      console.log("Speaker Submitted!", this.sessionForm.value);
      this.sessionForm.reset();
      this.sessions.push(newSession);
  }
}
