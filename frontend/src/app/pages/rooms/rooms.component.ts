import { Component, OnInit } from "@angular/core";
import { RoomService } from "../../services/room.service";
import { Room } from "../../data_models/room";
import { FormGroup, FormControl, Validators } from "@angular/forms";

@Component({
  selector: "app-rooms",
  templateUrl: "./rooms.component.html",
  styleUrls: ["./rooms.component.css"]
})
export class RoomsComponent implements OnInit {
  constructor(private roomService: RoomService) {}
  rooms: Room[];
  error: any;
  room = new Room("", 0);
  isEditable = false;
  roomForm = new FormGroup({
    roomName: new FormControl(""),
    roomCapacity: new FormControl("")
  });

  ngOnInit() {
    this.getAllRooms();
  }

  deleteRoom(id): void {
    if (confirm("Are you sure you want to remove it?")) {
      this.roomService
        .deleteRoom(id)
        .subscribe(error => (this.error = error));
      console.log("The following Room Deleted :", this.roomForm.value);
      this.rooms = this.rooms.filter(item => item.id !== id);
    }
  }

  getAllRooms(): void {
    this.roomService
      .getAllRooms()
      .subscribe(rooms => (this.rooms = rooms), error => (this.error = error));
  }

  onSubmit(): void {
    var newRoom = new Room(this.room.name, this.room.capacity);

    this.roomService
      .writeRoom(this.room.name, this.room.capacity)
      .subscribe(error => (this.error = error), id => (newRoom.id = id));
    console.log("Room Submitted!", this.roomForm.value);
    this.roomForm.reset();

    this.rooms.push(newRoom);
  }
}
