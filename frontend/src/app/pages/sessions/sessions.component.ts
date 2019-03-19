import { Component, OnInit, NgModule } from "@angular/core";

import { SessionService } from "../../services/session.service";
import { RoomService } from "../../services/room.service";
import { SpeakerService } from "../../services/speaker.service";
import { TimeslotService } from "../../services/timeslot.service";

import { Session } from "../../data_models/session";
import { Room } from "../../data_models/room";
import { Speaker } from "../../data_models/speaker";
import { Timeslot } from "../../data_models/timeslot";

import { FormGroup, FormControl, Validators } from "@angular/forms";

@Component({
  selector: "app-sessions",
  templateUrl: "./sessions.component.html",
  styleUrls: ["./sessions.component.css"]
})
export class SessionsComponent implements OnInit {
  constructor(
    private sessionService: SessionService,
    private roomService: RoomService,
    private speakerService: SpeakerService,
    private timeslotService: TimeslotService
  ) {}

  sessions: Session[];
  error: any;
  rooms: Room[];
  speakers: Speaker[];
  timeslots: Timeslot[];
  session = new Session(
    "",
    new Room("", 0),
    new Speaker("", "", ""),
    new Timeslot("", "")
  );
  currentSession = new Session(
    "",
    new Room("", 0),
    new Speaker("", "", ""),
    new Timeslot("", "")
  );
  selectedSession: Session;

  sessionForm = new FormGroup({
    sessionName: new FormControl(""),
    sessionRoom: new FormControl(""),
    sessionSpeaker: new FormControl(""),
    sessionTimeslot: new FormControl("")
  });

  ngOnInit() {
    this.getAllSessions();
    this.getAllRooms();
    this.getAllSpeakers();
    this.getAllTimeslots();
  }

  getAllSessions(): void {
    this.sessionService
      .getAllSessions()
      .subscribe(
        sessions => (this.sessions = sessions),
        error => (this.error = error)
      );
  }

  getAllRooms(): void {
    this.roomService
      .getAllRooms()
      .subscribe(rooms => (this.rooms = rooms), error => (this.error = error));
  }

  getAllSpeakers(): void {
    this.speakerService
      .getAllSpeakers()
      .subscribe(
        speakers => (this.speakers = speakers),
        error => (this.error = error)
      );
  }

  getAllTimeslots(): void {
    this.timeslotService
      .getAllTimeslots()
      .subscribe(
        timeslots => (this.timeslots = timeslots),
        error => (this.error = error)
      );
  }

  getSession(id: number) {
    this.sessionService.getSession(id);
  }

  updateSession(updatedSession: Session): void {
    if (confirm("Are you sure you want to update?")) {
      this.sessionService
        .updateSession(this.currentSession)
        .subscribe(error => (this.error = error));
      console.log("The following Session Udpated :", this.sessionForm.value);
      this.sessions = this.sessions.filter(item => item !== updatedSession);
      this.sessionForm.reset();
      this.sessions.push(updatedSession);
    }
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
      .subscribe(
        response => (newSession.id = response.id),
        error => (this.error = error)
      );
    console.log("Session Submitted!", this.sessionForm.value);
    this.sessionForm.reset();
    this.session.isEditable = false;
    this.sessions.push(newSession);
  }

  showEdit(session: Session): void {
    session.isEditable = true;
    this.currentSession.id = session.id;
  }

  cancel(session: Session): void {
    session.isEditable = false;
  }
}
