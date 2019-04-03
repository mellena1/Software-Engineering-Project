import { Component, OnInit, ViewChild } from "@angular/core";
import { RoomService } from "../../services/room.service";
import { Room } from "../../data_models/room";

import { Ng2SmartTableComponent } from "ng2-smart-table/ng2-smart-table.component";

import { TableSetting } from "../table_setting";
import { NumberInputComponent, TextRenderComponent, TextInputComponent } from "../../shared_components"
import { LocalDataSource } from "ng2-smart-table";

@Component({
  selector: "app-rooms",
  templateUrl: "./rooms.component.html"
})
export class RoomsComponent implements OnInit {
  constructor(private roomService: RoomService) {
    this.tableDataSource = new LocalDataSource();

    this.roomService.getAllRooms().subscribe(data => {
      this.tableDataSource.load(data);
    })
  }

  rooms: Room[];
  @ViewChild("table") table: Ng2SmartTableComponent;

  columns = {
    name: {
      title: "Room Name",
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "custom",
        component: TextInputComponent
      },
      filter: false
    },
    capacity: {
      title: "Room Capacity",
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "custom",
        component: NumberInputComponent
      },
      filter: false
    }
  };
  tableSettings = new TableSetting(this.columns);
  tableDataSource: LocalDataSource;

  ngOnInit() {
    this.getAllRooms();

    this.table.userRowSelect.subscribe(_ => {
      this.table.grid.dataSet.deselectAll();
    });

    this.table.initGrid();
    this.table.grid.createFormShown = true;
  }

  getAllRooms(): void {
    this.roomService
      .getAllRooms()
      .subscribe(rooms => (this.rooms = rooms), error => console.log(error));
  }

  deleteRoom(event): void {
    var id = event.data.id;

    this.roomService.deleteRoom(id).subscribe(_ => _, error => console.log(error));
    event.confirm.resolve();
    this.rooms = this.rooms.filter(item => item.id !== id);
    event.source.refresh();
  }

  addARoom(event): void {
    var room = event.newData;
    this.roomService.writeRoom(room.name, room.capacity).subscribe(
      response => {
        console.log("hi");
        event.source.data[0].id = response.id;
        console.log(event.source);
        event.confirm.resolve();
        this.table.create.emit();
      },
      error => {
        console.log(error);
        event.confirm.reject();
      }
    );
  }

  updateRoom(event): void {
    var room = event.newData;

    this.roomService.updateRoom(room).subscribe(
      _ => {
        event.confirm.resolve();
      },
      error => {
        event.confirm.reject();
        console.log(error);
      }
    );
  }
}
