import { Component, OnInit, ViewChild } from "@angular/core";
import { SessionService } from "../../services/session.service";

import { Ng2SmartTableComponent } from "ng2-smart-table/ng2-smart-table.component";

import { TableSetting } from "../table_setting";
import {
  TextRenderComponent,
  TextInputComponent,
  TextListInputComponent,
  TimeslotRenderComponent,
  TimeslotListInputComponent
} from "../../shared_components";
import { LocalDataSource } from "ng2-smart-table";
import { Session } from "src/app/data_models/session";
import { Speaker } from "src/app/data_models/speaker";
import { Room } from "src/app/data_models/room";
import { Timeslot } from "src/app/data_models/timeslot";
import { RoomService } from "src/app/services/room.service";
import { SpeakerService } from "src/app/services/speaker.service";
import { TimeslotService } from "src/app/services/timeslot.service";
import { ErrorGlobals } from "src/app/globals/errors.global";
import { TimeslotGlobals } from "src/app/globals/timeslot.global";

@Component({
  selector: "app-sessions",
  templateUrl: "./sessions.component.html"
})
export class SessionsComponent implements OnInit {
  @ViewChild("table") table: Ng2SmartTableComponent;
  tableDataSource: LocalDataSource;
  roomsListForTable: object[] = [{ value: null, title: "" }];
  speakersListForTable: object[] = [{ value: null, title: "" }];
  timeslotsListForTable: object[] = [{ value: null, title: "" }];

  columns = {
    name: {
      title: "Session Name",
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "custom",
        component: TextInputComponent
      }
    },
    room: {
      title: "Room Name",
      valuePrepareFunction: (room: Room) => {
        if (room === null) {
          return "";
        } else {
          return room.name;
        }
      },
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "custom",
        component: TextListInputComponent,
        config: {
          list: this.roomsListForTable
        }
      }
    },
    speaker: {
      title: "Speaker Name",
      valuePrepareFunction: (speaker: Speaker) => {
        if (speaker != null && typeof speaker === "object") {
          return Speaker.getFullName(speaker);
        } else {
          return "";
        }
      },
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "custom",
        component: TextListInputComponent,
        config: {
          list: this.speakersListForTable
        }
      }
    },
    timeslot: {
      title: "Timeslot",
      valuePrepareFunction: (timeslot: Timeslot) => {
        if (timeslot == null) {
          return "";
        } else {
          return timeslot.startTime + " " + timeslot.endTime;
        }
      },
      type: "custom",
      renderComponent: TimeslotRenderComponent,
      editor: {
        type: "custom",
        component: TimeslotListInputComponent,
        config: {
          list: this.timeslotsListForTable
        }
      }
    }
  };
  tableSettings: TableSetting = new TableSetting(this.columns);

  constructor(
    private sessionService: SessionService,
    private roomService: RoomService,
    private speakerService: SpeakerService,
    private timeslotService: TimeslotService,
    private timeslotGlobals: TimeslotGlobals,
    private errorGlobals: ErrorGlobals
  ) {
    this.tableDataSource = new LocalDataSource();

    this.getAllSessions();
    this.getAllRooms();
    this.getAllSpeakers();
    this.getAllTimeslots();
  }

  ngOnInit() {
    this.table.userRowSelect.subscribe(() => {
      this.table.grid.dataSet.deselectAll();
    });

    this.table.initGrid();

    this.tableDataSource.onChanged().subscribe(() => {
      this.table.grid.createFormShown = true;
    });
  }

  getAllSessions() {
    this.sessionService.getAllSessions().subscribe(
      (data: Session[]) => {
        data.sort((a, b) => {
          return a.name < b.name ? -1 : 1;
        });
        this.tableDataSource.load(data);
      },
      error => {
        console.log(error);
      }
    );
  }

  getAllRooms() {
    this.roomService.getAllRooms().subscribe(
      (data: Room[]) => {
        data.sort((a, b) => {
          return a.name < b.name ? -1 : 1;
        });
        data.forEach(room => {
          this.roomsListForTable.push({
            value: room,
            title: room.name
          });
        });
        this.tableSettings = new TableSetting(this.columns);
      },
      error => {
        console.log(error);
      }
    );
  }

  getAllSpeakers() {
    this.speakerService.getAllSpeakers().subscribe(
      (data: Speaker[]) => {
        data.sort((a, b) => {
          return a.firstName < b.firstName ? -1 : 1;
        });
        data.forEach(speaker => {
          this.speakersListForTable.push({
            value: speaker,
            title: `${speaker.firstName} ${speaker.lastName}`
          });
        });
        this.tableSettings = new TableSetting(this.columns);
      },
      error => {
        console.log(error);
      }
    );
  }

  getAllTimeslots() {
    this.timeslotService.getAllTimeslots().subscribe(
      (data: Timeslot[]) => {
        data.sort((a, b) => {
          return TimeslotGlobals.sortTime(a.startTime, b.startTime);
        });
        data.forEach(timeslot => {
          this.timeslotsListForTable.push({
            value: timeslot,
            title: timeslot.startTime + " " + timeslot.endTime
          });
        });
        this.tableSettings = new TableSetting(this.columns);
      },
      error => {
        console.log(error);
      }
    );
  }

  addASession(event): void {
    var session = event.newData;
    this.sessionService
      .writeSession(
        session.name,
        session.room.id,
        session.speaker.id,
        session.timeslot.id
      )
      .subscribe(
        response => {
          session.id = response.id;
          event.confirm.resolve(session);
        },
        error => {
          console.log(error);
          if (error.status === 503) {
            this.errorGlobals.newError(
              "The server is unavailable, please wait a minute and try again"
            );
          } else {
            this.errorGlobals.newError(
              "You must set one of the fields to add a session"
            );
          }
          event.confirm.reject();
        }
      );
  }

  updateSession(event): void {
    var session = event.newData;
    this.sessionService.updateSession(session).subscribe(
      () => {
        event.confirm.resolve(session);
      },
      error => {
        console.log(error);
        if (error.status === 503) {
          this.errorGlobals.newError(
            "The server is unavailable, please wait a minute and try again"
          );
        } else {
          this.errorGlobals.newError(
            "You must change one of the values and set at least of the fields"
          );
        }
        event.confirm.reject();
      }
    );
  }

  deleteSession(event): void {
    this.sessionService.deleteSession(event.data.id).subscribe(
      () => {},
      error => {
        console.log(error);
        if (error.status === 503) {
          this.errorGlobals.newError(
            "The server is unavailable, please wait a minute and try again"
          );
        } else {
          this.errorGlobals.newError("Delete failed");
        }
        event.confirm.reject();
      }
    );
    event.confirm.resolve();
  }
}
