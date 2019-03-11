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
  newRoom: Room;
  selectedRoom: Room;
  error: any;
  public show: boolean = false;
  public buttonName: any = "Add a Room"
  roomForm: FormGroup;

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

  onSelect(room: Room): void {
    this.selectedRoom = room;
  }

  toggle() {
    this.show = !this.show;
    if (this.show) {
      this.buttonName = "Hide";
    } else {
      this.buttonName = "Add a Room";
    }
  }
}