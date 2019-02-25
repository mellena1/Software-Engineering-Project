import { Component, OnInit } from '@angular/core';

import { RoomService } from '../services/room.service'
import { Room } from '../data_models/room';

@Component({
  selector: 'app-rooms',
  templateUrl: './rooms.component.html',
  styleUrls: ['./rooms.component.css']
})
export class RoomsComponent implements OnInit {
  // rooms: Room[] = [
  //   { id: 1, capacity: 10, roomName: "test" },
  //   { id: 1, capacity: 10, roomName: "dog" },
  //   { id: 1, capacity: 10, roomName: "ayyy" }
  // ];

  rooms: Room[]
  constructor(private roomService: RoomService ) { }
  
  ngOnInit() {
    this.getAllRooms();
  }

  getAllRooms(): void {
    this.roomService.getAllRooms()
    .subscribe(rooms => (this.rooms = rooms))
  }
}
