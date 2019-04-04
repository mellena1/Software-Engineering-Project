import { Component, OnInit, ViewChild } from "@angular/core";
import { SessionService } from "../../services/session.service";

import { Ng2SmartTableComponent } from "ng2-smart-table/ng2-smart-table.component";

import { TableSetting } from "../table_setting";
import {
  TextRenderComponent,
  TextInputComponent
} from "../../shared_components";
import { LocalDataSource } from "ng2-smart-table";
import { Speaker } from "src/app/data_models/speaker";
import { Room } from "src/app/data_models/room";

@Component({
  selector: "app-sessions",
  templateUrl: "./sessions.component.html"
})
export class SessionsComponent implements OnInit {
  @ViewChild("table") table: Ng2SmartTableComponent;
  tableDataSource: LocalDataSource;

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
      valuePrepareFunction: (room: Room) => { return room.name },
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "custom",
        component: TextInputComponent
      },
    },
    speaker: {
      title: "Speaker Name",
      valuePrepareFunction: (speaker: Speaker) => { return speaker.lastName },
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "custom",
        component: TextInputComponent
      }
    },
    timeslot: {
      title: "Timeslot",
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "custom",
        component: TextInputComponent
      }
    }
  };
  tableSettings = new TableSetting(this.columns);

  constructor(private sessionService: SessionService) {
    this.tableDataSource = new LocalDataSource();

    this.sessionService.getAllSessions().subscribe(data => {
      this.tableDataSource.load(data);
    });
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

  addASession(event): void {
    var session = event.newData;
    this.sessionService
      .writeSession(session.name, session.room.id, session.speaker.id, session.timeslot.id)
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
      () => { },
      error => {
        console.log(error);
      }
    );
    event.confirm.resolve();
  }
}
