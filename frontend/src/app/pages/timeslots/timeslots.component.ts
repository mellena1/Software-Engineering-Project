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
  selectedTimeslot = new Timeslot(1, "", "");
  error: any;

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
  }

  updateTimeslot(): void {
    this.timeslotService
    .updateTimeslot()
  }

  onSelect(timeslot: Timeslot): void {
    this.selectedTimeslot = timeslot;
    this.addTimeslot(this.selectedTimeslot);
  }

}
