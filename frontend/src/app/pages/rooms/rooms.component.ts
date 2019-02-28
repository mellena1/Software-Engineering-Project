import { Component, OnInit } from '@angular/core';

import { RoomService } from '../../services/room.service'
import { Room } from '../../data_models/room';

@Component({
  selector: 'app-rooms',
  templateUrl: './rooms.component.html',
  styleUrls: ['./rooms.component.css']
})
export class RoomsComponent implements OnInit {
  rooms: Room[];
  selectedRoom: Room;
  error: any;
  public show:boolean = false;
  public buttonName:any = "Add a Room"
  constructor(private roomService: RoomService ) { }
  
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
  
  addRoom(): void {
      this.roomService
      .writeRoom()
      this.show=false;
  }
  
  updateRoom(): void{
    this.roomService
    .updateRoom()
  }

  deleteRoom(): void{
    this.roomService
    .deleteRoom()
  }

  onSelect(room: Room): void {
    this.selectedRoom = room;
  }

  toggle(){
    this.show = !this.show;
  

  if(this.show){
    this.buttonName = "Hide";
  }
  else{
    this.buttonName = "Add a Room";
  }
}

}
