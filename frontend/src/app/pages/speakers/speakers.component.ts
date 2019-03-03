import { Component, OnInit } from '@angular/core';
import {SpeakerService} from 'src/app/services/speaker.service'
import {Speaker} from 'src/app/data_models/speaker';
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
  error: any;

  ngOnInit() {
    this.getAllSpeakers();
    this.getSpeaker(1);
    console.log('deleting speaker...')
    this.speakerService.deleteSpeaker(1).subscribe(data => console.log(data));
    console.log('deleted speaker...')
  }

  getAllSpeakers(): void {
    this.speakerService
      .getAllSpeakers()
      .subscribe(
        speakers => (this.speakers = speakers),
        error   => (this.error = error)
      )
  }

  getSpeaker(id: number): void {
    this.speakerService
      .getSpeaker(id)
      .subscribe(
        speaker => (this.speaker = speaker),
        error => (this.error = error)
      )
  }

  writeSpeaker(speaker: Speaker): void {
    this.speakerService.writeSpeaker(speaker)
  }
}
