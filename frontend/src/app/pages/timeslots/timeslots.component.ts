import { Component, OnInit } from '@angular/core';
import { Timeslot } from 'src/app/data_models/timeslot';
import { TimeslotService } from 'src/app/services/timeslot.service';

@Component({
  selector: 'app-timeslots',
  templateUrl: './timeslots.component.html',
  styleUrls: ['./timeslots.component.css']
})
export class TimeslotsComponent implements OnInit {
  constructor(private timeslotService: TimeslotService) { }
  timeslots: Timeslot[];
  timeslot: Timeslot;
  error: any;

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
}
