import { Component, OnInit } from "@angular/core";
import { SpeakerService } from "src/app/services/speaker.service";
import { Speaker } from "src/app/data_models/speaker";
import { FormGroup, FormControl, Validators } from "@angular/forms";

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
  isEditable = false;
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
    if (confirm("Are you sure you want to remove it?")) {
      this.speakerService
        .deleteSpeaker(id)
        .subscribe(error => (this.error = error));
      console.log("The following Speaker Deleted :", this.speakerForm.value);
      this.speakers = this.speakers.filter(item => item.id !== id);
    }
  }

  onSubmit(): void {
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
        response => (newSpeaker.id = response.id),
        error => (this.error = error)
      );
    console.log("Speaker Submitted!", this.speakerForm.value);
    this.speakerForm.reset();
    this.speakers.push(newSpeaker);
  }
}
