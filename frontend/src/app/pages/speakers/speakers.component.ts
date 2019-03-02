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
    this.speakerService.getSpeaker(1);
  }

  getAllSpeakers(): void {
    this.speakerService
      .getAllSpeakers()
      .subscribe(
        speakers => (this.speakers = speakers),
        error   => (this.error = error)
      )
  }
}
