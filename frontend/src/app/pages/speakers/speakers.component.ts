import { Component, OnInit } from "@angular/core";
import { FormControl, FormGroup, Validators } from "@angular/forms";
import { Speaker } from "src/app/data_models/speaker";
import { SpeakerService } from "src/app/services/speaker.service";

@Component({
  selector: "app-speakers",
  templateUrl: "./speakers.component.html",
  styleUrls: ["./speakers.component.css"]
})
export class SpeakersComponent implements OnInit {
  constructor(private speakerService: SpeakerService) {}
  speakers: Speaker[];
  error: any;
  speaker = new Speaker("", "", "");
  currentSpeaker = new Speaker("", "", "");
  //isEditable = false;
  speakerForm = new FormGroup({
    firstName: new FormControl(""),
    lastName: new FormControl(""),
    email: new FormControl("")
  });

  ngOnInit() {
    this.getAllSpeakers();
  }

  getAllSpeakers(): void {
    this.speakerService
      .getAllSpeakers()
      .subscribe(
        speakers => (this.speakers = speakers),
        error => (this.error = error)
      );
  }

  deleteSpeaker(id): void {
    this.speakerService.deleteSpeaker(id).subscribe(error => {
      this.error = error;
      if (this.error == null) {
        this.speakers = this.speakers.filter(speaker => speaker.id !== id);
        console.log("The following Speaker Deleted :", this.speakerForm.value);
      }
    });
  }

  writeSpeaker(): void {
    var newSpeaker = new Speaker(
      this.speaker.firstName,
      this.speaker.lastName,
      this.speaker.email
    );
    this.speakerService
      .writeSpeaker(
        this.speaker.firstName,
        this.speaker.lastName,
        this.speaker.email
      )
      .subscribe(
        response => {
          newSpeaker.id = response.id;
          this.speakerForm.reset();
          this.speakers.push(newSpeaker);
          console.log("Speaker Submitted!", this.speakerForm.value);
        },
        error => {
          this.error = error;
          console.log(this.error);
        }
      );
  }

  updateSpeaker(): void {
    var index = this.speakers.findIndex(
      item => item.id === this.currentSpeaker.id
    );
    var updatedSpeaker = this.speakers[index];

    this.speakers[index].isEditable = false;
    this.speakerService
      .updateSpeaker(this.currentSpeaker)
      .subscribe(error => {this.error = error;
    });

    console.log("The following Speaker Udpated :", this.speakerForm.value);


    updatedSpeaker.firstName = this.currentSpeaker.firstName;
    updatedSpeaker.lastName = this.currentSpeaker.lastName;
    updatedSpeaker.email = this.currentSpeaker.email;

    this.speakerForm.reset();
  }

  showEdit(speaker: Speaker): void {
    speaker.isEditable = true;
    this.currentSpeaker.id = speaker.id;
    this.currentSpeaker.firstName = speaker.firstName;
    this.currentSpeaker.lastName = speaker.lastName;
    this.currentSpeaker.email = speaker.email;
  }

  cancel(speaker: Speaker): void {
    speaker.isEditable = false;
    this.speakerForm.reset();
  }
}
