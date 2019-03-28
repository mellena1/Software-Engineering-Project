import { Component, OnInit } from "@angular/core";
import { RoomService } from "../../services/room.service";
import { Room } from "../../data_models/room";

import { TableSetting } from '../table_setting';
import { NumberInputComponent } from '../../shared_components/number_input.component'

@Component({
  selector: "app-rooms",
  templateUrl: "./rooms.component.html",
  styleUrls: ["./rooms.component.css"]
})
export class RoomsComponent implements OnInit {
  constructor(private roomService: RoomService) { }
  rooms: Room[];
  error: any;
  tableSettings: object;

  ngOnInit() {
    var columns = {
      name: {
        title: 'Room Name'
      },
      capacity: {
        title: 'Room Capacity',
        editor: {
          type: 'custom',
          component: NumberInputComponent,
        },
      },
    };
    this.tableSettings = new TableSetting(columns);
    this.getAllRooms();
  }

  getAllRooms(): void {
    this.roomService
      .getAllRooms()
      .subscribe(rooms => (this.rooms = rooms), error => (this.error = error));
  }

  deleteRoom(event): void {
    var id = event.data.id;

    console.log(id);

    this.roomService.deleteRoom(id).subscribe(error => (this.error = error));
    event.confirm.resolve();
    this.rooms = this.rooms.filter(item => item.id !== id);
    event.source.refresh();
  }

  addARoom(event): void {
    var room = event.newData;
    room.capacity = +room.capacity; // convert to number
    this.roomService
      .writeRoom(room.name, room.capacity)
      .subscribe(
        response => {
          room.id = response.id
          this.rooms.push(room);
          event.confirm.resolve();
        },
        error => {
          this.error = error
          event.confirm.reject();
        }
      );
  }

  updateRoom(event): void {
    var room = event.newData;

    this.roomService.updateRoom(room).subscribe(_ =>{
        event.confirm.resolve();
      }, 
      error => {
        event.confirm.reject();
        this.error = error;
        console.log(error);
    });
  }
}
