import { Component, OnInit } from '@angular/core';
import { RoomService } from '../../services/room.service'
import { Room } from '../../data_models/room';
import { FormGroup, FormControl, Validators } from '@angular/forms';

@Component({
  selector: 'app-rooms',
  templateUrl: './rooms.component.html',
  styleUrls: ['./rooms.component.css']
})
export class RoomsComponent implements OnInit {
  constructor(private roomService: RoomService) { }
  rooms: Room[];
  room: Room;
  error: any;
  
  ngOnInit() {
    this.getAllRooms();
  }

  getAllRooms(): void {
    this.roomService
      .getAllRooms()
      .subscribe(
        rooms => (this.rooms = rooms),
        error => (this.error = error)
      )
  }
}
