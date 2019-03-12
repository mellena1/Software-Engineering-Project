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
  error: any;
  room={
    name:"",
    capacity:0
  }
  isEditable = false;
  roomForm = new FormGroup({
    roomName: new FormControl(''),
    roomCapacity: new FormControl(''),
  });



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
  
  editRoom() {
    this.isEditable = true;
  }

  onSubmit(): void{
    this.roomService
      .writeRoom(this.room.name,this.room.capacity)
      .subscribe(
        error => (this.error = error)
      )
    console.log("Room Submitted!", this.roomForm.value);
    this.roomForm.reset();
    window.location.reload();
  }
}