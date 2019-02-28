import { Component, OnInit } from '@angular/core';
import { Timeslot } from 'src/app/data_models/timeslot';
import { TimeslotService } from 'src/app/services/timeslot.service';

@Component({
  selector: 'app-timeslots',
  templateUrl: './timeslots.component.html',
  styleUrls: ['./timeslots.component.css']
})
export class TimeslotsComponent implements OnInit {
  timeslots: Timeslot[];
  selectedTimeslot: Timeslot;
  error: any;
  public show:boolean = false;
  public buttonName:any = "Add a Timeslot"

  constructor(private timeslotService: TimeslotService) { }

  ngOnInit() {
    this.getAllTimeslots();
  }

  getAllTimeslots(): void{
    this.timeslotService
      .getAllTimeslots()
      .subscribe(
        timeslots => (this.timeslots = timeslots),
        error => (this.error = error)
       )
  }

  addTimeslot(timeslot: Timeslot): void {
    this.timeslotService
    .writeTimeslot(timeslot)
    this.show=false;
  }

  updateTimeslot(updatedTimeslot: Timeslot): void {
    this.timeslotService
    .updateTimeslot(updatedTimeslot)
  }

  onSelect(timeslot: Timeslot): void {
    this.selectedTimeslot = timeslot;
    this.addTimeslot(this.selectedTimeslot);
  }

  toggle(){
    this.show = !this.show;
    if(this.show){
      this.buttonName = "Hide";
    }
    else{
      this.buttonName = "Add a Timeslot";
    }
  }
}
