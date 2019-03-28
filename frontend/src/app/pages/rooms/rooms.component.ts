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
  currentRoom = new Room("", 0);
  emptyRoom = new Room ("", 0);
  disableEdit= false;


  roomForm = new FormGroup({
    roomName: new FormControl(""),
    roomCapacity: new FormControl("")
  });


  ngOnInit() {
    this.getAllRooms();
  }

  getAllRooms(): void {
    this.roomService
      .getAllRooms()
      .subscribe(rooms => (this.rooms = rooms), error => (this.error = error));
  }

  deleteRoom(id): void {
    this.roomService.deleteRoom(id).subscribe(error => (this.error = error));
    this.rooms = this.rooms.filter(item => item.id !== id);
  }

  onSubmit(): void {
    var newRoom = new Room(this.room.name, this.room.capacity);
    this.roomService
      .writeRoom(this.room.name, this.room.capacity)
      .subscribe(
        response => (newRoom.id = response.id),
        error => (this.error = error)
      );
    this.roomForm.reset();
    this.rooms.push(newRoom);
  }

  updateRoom(): void {
    var index = this.rooms.findIndex(
      item => item.id === this.currentRoom.id
    );
    var updatedRoom = this.rooms[index];

    this.rooms[index].isEditable = false;
    this.roomService
      .updateRoom(this.currentRoom)
      .subscribe(error => {this.error = error;
    });

    console.log("The following Room Udpated :", this.roomForm.value);


    updatedRoom.name = this.currentRoom.name;
    updatedRoom.capacity = this.currentRoom.capacity;

    this.disableEdit = false;
  }

  showEdit(room: Room): void {
    room.isEditable = true;
    this.currentRoom = room;
    this.disableEdit = true;
  }

  cancel(room: Room): void {
    room.isEditable = false;
    this.currentRoom = this.emptyRoom;
    this.disableEdit = false;
  }
}
