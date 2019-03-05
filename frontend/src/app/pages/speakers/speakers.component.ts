import { Component, OnInit } from '@angular/core';
import { SpeakerService } from 'src/app/services/speaker.service'
import { Speaker } from 'src/app/data_models/speaker';
import { FormGroup, FormControl, Validators } from '@angular/forms';


@Component({
  selector: 'app-speakers',
  templateUrl: './speakers.component.html',
  styleUrls: ['./speakers.component.css']
})
export class SpeakersComponent implements OnInit {
  constructor(private speakerService: SpeakerService) { }
  speaker: Speaker;
  speakers: Speaker[];
  selectedSpeaker: Speaker;
  error: any;
  public show: boolean = false;
  public buttonName: any = "Add a Speaker"
  speakerForm: FormGroup;

  ngOnInit() {
    this.getAllSpeakers();
    this.speakerForm = new FormGroup({
      'firstName': new FormControl(this.speaker.firstName, [
        Validators.required
      ]),
      'lastName': new FormControl(this.speaker.lastName, [
        Validators.required
      ]),
      'email': new FormControl(this.speaker.email, [
        Validators.required
      ])
    });
  }

  getAllSpeakers(): void {
    this.speakerService
      .getAllSpeakers()
      .subscribe(
        speakers => (this.speakers = speakers),
        error => (this.error = error)
      )
  }

  onSelect(speaker: Speaker): void {
    this.selectedSpeaker = speaker;
  }

  toggle() {
    this.show = !this.show;
    if (this.show) {
      this.buttonName = "Hide";
    }
    else {
      this.buttonName = "Add a Room";
    }
  }
}