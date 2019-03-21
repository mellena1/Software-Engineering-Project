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

  //  isEditable = false;
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

  deleteSpeaker(id): void {
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
    var index = this.rooms.findIndex(item => item.id === this.currentRoom.id);
    var currentSpeaker = this.room[index];

    currentSpeaker.isEditable = false;
    this.roomService.updateRoom(this.currentRoom).subscribe(error => {
      this.error = error;
    });

    currentSpeaker.name = this.currentRoom.name;
    currentSpeaker.capacity = this.currentRoom.capacity;

    this.roomForm.reset();
  }

  showEdit(room: Room): void {
    room.isEditable = true;
    this.currentRoom.id = room.id;
  }

  cancel(room: Room): void {
    room.isEditable = false;
    this.roomForm.reset();
  }
}
