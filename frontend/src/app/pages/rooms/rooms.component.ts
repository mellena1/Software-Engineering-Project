import { Component, OnInit, ViewChild } from "@angular/core";
import { RoomService } from "../../services/room.service";

import { Ng2SmartTableComponent } from "ng2-smart-table/ng2-smart-table.component";

import { TableSetting } from "../table_setting";
import {
  NumberInputComponent,
  TextRenderComponent,
  TextInputComponent
} from "../../shared_components";
import { LocalDataSource } from "ng2-smart-table";
import { ErrorGlobals } from "../../globals/errors.global";
import { HttpErrorResponse } from "@angular/common/http";

@Component({
  selector: "app-rooms",
  templateUrl: "./rooms.component.html"
})
export class RoomsComponent implements OnInit {
  @ViewChild("table") table: Ng2SmartTableComponent;
  tableDataSource: LocalDataSource;

  columns = {
    name: {
      title: "Room Name",
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "custom",
        component: TextInputComponent
      }
    },
    capacity: {
      title: "Room Capacity",
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "custom",
        component: NumberInputComponent
      }
    }
  };
  tableSettings = new TableSetting(this.columns);

  constructor(
    private roomService: RoomService,
    private errorGlobals: ErrorGlobals
  ) {
    this.tableDataSource = new LocalDataSource();

    this.getAllRooms();
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
      data => {
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

  addARoom(event): void {
    var room = event.newData;
    this.roomService.writeRoom(room.name, room.capacity).subscribe(
      response => {
        room.id = response.id;
        event.confirm.resolve(room);
      },
      (error: HttpErrorResponse) => {
        if (error.status === 503) {
          this.errorGlobals.newError(
            "The server is unavailable, please wait a minute and try again"
          );
        } else {
          this.errorGlobals.newError("Must provide a room name");
        }
        console.log(error);
        event.confirm.reject();
      }
    );
  }

  updateRoom(event): void {
    var room = event.newData;
    this.roomService.updateRoom(room).subscribe(
      () => {
        event.confirm.resolve(room);
      },
      error => {
        if (error.status === 503) {
          this.errorGlobals.newError(
            "The server is unavailable, please wait a minute and try again"
          );
        } else {
          this.errorGlobals.newError(
            "Must change a value and room name must be set"
          );
        }
        console.log(error);
        event.confirm.reject();
      }
    );
  }

  deleteRoom(event): void {
    this.roomService.deleteRoom(event.data.id).subscribe(
      () => {
        event.confirm.resolve();
      },
      error => {
        if (error.status === 503) {
          this.errorGlobals.newError(
            "The server is unavailable, please wait a minute and try again"
          );
        } else {
          this.errorGlobals.newError("Delete failed");
        }
        console.log(error);
        event.confirm.reject();
      }
    );
  }
}
