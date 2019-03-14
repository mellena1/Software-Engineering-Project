import { Component, OnInit } from '@angular/core';
import { Timeslot } from 'src/app/data_models/timeslot';
import { TimeslotService } from 'src/app/services/timeslot.service';
import { FormGroup, FormControl, Validators } from '@angular/forms';

@Component({
  selector: "app-timeslots",
  templateUrl: "./timeslots.component.html",
  styleUrls: ["./timeslots.component.css"]
})
export class TimeslotsComponent implements OnInit {
  constructor(private timeslotService: TimeslotService) {}
  timeslots: Timeslot[];
  error: any;
  timeslot = new Timeslot("", "");
  oldTimeslot = new Timeslot("", "");
  isEditable: boolean;
  timeslotForm = new FormGroup({
    timeslotStart: new FormControl(''),
    timeslotEnd: new FormControl(''),
  });

  ngOnInit() {
    this.getAllTimeslots();
    this.isEditable = false;
  }

  getAllTimeslots(): void {
    this.timeslotService
      .getAllTimeslots()
      .subscribe(
        timeslots => (this.timeslots = timeslots),
        error => (this.error = error)
      );
  }

  onSubmit(): void{
    var newTimeslot = new Timeslot(this.timeslot.startTime, this.timeslot.endTime);

    this.timeslotService
      .writeTimeslot(this.timeslot.startTime, this.timeslot.endTime)
      .subscribe(
        error => (this.error = error),
        id => (newTimeslot.id = id)
      )
    console.log("Timeslot Submitted!", this.timeslotForm.value);
    this.timeslotForm.reset();

    this.timeslots.push(newTimeslot);
  }

  deleteTimeslot(timeslotid): void {
    if(confirm("Are you sure you want to remove it?"))
    {
      this.timeslotService
      .deleteTimeslot(timeslotid)
      .subscribe(
        error => (this.error = error)
      )
      console.log("The following Timeslot Deleted :", this.timeslotForm.value);
      this.timeslots = this.timeslots.filter(item => item.id !== timeslotid);
    }
  }

  updateTimeslot(timeslot: Timeslot): void {
    if(confirm("Are you sure you want to update?"))
    {
      this.timeslotService
      .updateTimeslot(timeslot)
      .subscribe(
        error => (this.error=error)
      )
      console.log("The following Timeslot Udpated :", this.timeslotForm.value);
      this.timeslots = this.timeslots.filter(item => item !== timeslot);
    }
  }

  showEdit(timeslot: Timeslot): void {
    this.oldTimeslot = timeslot;
    timeslot.isEditable = true;
  }

  cancel(timeslot: Timeslot): void {
    timeslot = this.oldTimeslot;
    timeslot.isEditable = false;
  }

}
