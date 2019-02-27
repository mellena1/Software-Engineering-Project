import { Component, OnInit } from '@angular/core';
import {SpeakerService} from 'src/app/services/speaker.service'
import {Speaker} from 'src/app/data_models/speaker';


@Component({
  selector: 'app-speakers',
  templateUrl: './speakers.component.html',
  styleUrls: ['./speakers.component.css']
})
export class SpeakersComponent implements OnInit {
  speaker: Speaker[];
  selectedSpeaker: Speaker;
  error: any;
  constructor(private speakerService: SpeakerService) { }

  ngOnInit() {
    this.getAllSpeakers();
  }

  getAllSpeakers(): void {
    this.speakerService
      .getAllSpeakers()
      .subscribe(
        speaker => (this.speaker = speaker),
        error   => (this.error = error)
      )
  }

  onSelect(speaker: Speaker): void {
    this.selectedSpeaker = speaker;
  }

}
