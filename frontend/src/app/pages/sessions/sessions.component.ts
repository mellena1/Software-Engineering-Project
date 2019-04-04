import { Component, OnInit, ViewChild } from "@angular/core";
import { SessionService } from "../../services/session.service";

import { Ng2SmartTableComponent } from "ng2-smart-table/ng2-smart-table.component";

import { TableSetting } from "../table_setting";
import {
  TextRenderComponent,
  TextInputComponent,
  TimeslotRenderComponent
} from "../../shared_components";
import { LocalDataSource } from "ng2-smart-table";
import { Speaker } from "src/app/data_models/speaker";
import { Room } from "src/app/data_models/room";
import { Timeslot } from "src/app/data_models/timeslot";
import { RoomService } from "src/app/services/room.service";
import { SpeakerService } from "src/app/services/speaker.service";
import { TimeslotService } from "src/app/services/timeslot.service";

import { TimeslotGlobals } from "../../globals/timeslot.global";

@Component({
  selector: "app-sessions",
  templateUrl: "./sessions.component.html"
})
export class SessionsComponent implements OnInit {
  @ViewChild("table") table: Ng2SmartTableComponent;
  tableDataSource: LocalDataSource;
  roomsListForTable: object[] = [{ value: null, title: "-" }];
  speakersListForTable: object[] = [{ value: null, title: "-" }];
  timeslotsListForTable: object[] = [{ value: null, title: "-" }];

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
      valuePrepareFunction: room => {
        if (room === null) {
          return "";
        } else {
          return room.name;
        }
      },
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "list",
        config: {
          list: this.roomsListForTable
        }
      }
    },
    speaker: {
      title: "Speaker Name",
      valuePrepareFunction: speaker => {
        if (speaker == null) {
          return "";
        } else {
          return speaker.firstName + " " + speaker.lastName;
        }
      },
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "list",
        config: {
          list: this.speakersListForTable
        }
      }
    },
    timeslot: {
      title: "Timeslot",
      valuePrepareFunction: timeslot => {
        if (timeslot == null) {
          return "";
        } else {
          return timeslot.startTime + " " + timeslot.endTime;
        }
      },
      type: "custom",
      renderComponent: TimeslotRenderComponent,
      editor: {
        type: "list",
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
    private timeslotGlobals: TimeslotGlobals
  ) {
    this.tableDataSource = new LocalDataSource();

    this.sessionService.getAllSessions().subscribe(
      data => {
        this.tableDataSource.load(data);
      },
      error => {
        console.log(error);
      }
    );

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

  getAllRooms() {
    this.roomService.getAllRooms().subscribe(
      (data: Room[]) => {
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
        data.forEach(timeslot => {
          this.timeslotsListForTable.push({
            value: timeslot,
            title: `${this.timeslotGlobals.formatTime(timeslot.startTime)}-${this.timeslotGlobals.formatTime(timeslot.endTime)}`
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
        event.confirm.reject();
      }
    );
  }

  deleteSession(event): void {
    this.sessionService.deleteSession(event.data.id).subscribe(
      () => {},
      error => {
        console.log(error);
      }
    );
    event.confirm.resolve();
  }
}
