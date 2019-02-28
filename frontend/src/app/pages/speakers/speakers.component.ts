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
  public show:boolean = false;
  public buttonName:any = "Add a Speaker"
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

   
  addSpeaker(): void {
    this.speakerService
    .writeSpeaker()
    this.show=false;
}

updateSpeaker(): void{
  this.speakerService
  .updateSpeaker
}

deleteRoom(): void{
  this.speakerService
  .deleteSpeaker
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

