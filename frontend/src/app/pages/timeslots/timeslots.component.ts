import { Component, OnInit } from "@angular/core";
import { Timeslot } from "src/app/data_models/timeslot";
import { TimeslotService } from "src/app/services/timeslot.service";
import { FormGroup, FormControl, Validators } from "@angular/forms";

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

  eventDate = "2019-04-06";
  twelveHourIsChecked = true;
  timeFormat = "12hour";
  
  seconds = ":00";
  startHour = "00";
  startMin = "00";
  endHour = "00";
  endMin = "00";
  
  timeslotForm = new FormGroup({
    timeslotStart: new FormControl(""),
    timeslotEnd: new FormControl("")
  });

  ngOnInit() {
    this.getAllTimeslots();
  }

  getAllTimeslots(): void {
    this.timeslotService
      .getAllTimeslots()
      .subscribe(
        timeslots => (this.timeslots = timeslots),
        error => (this.error = error)
      );
  }
  
  writeTimeslot(startTime: string, endTime: string): number {
    var newId: number
    this.timeslotService
      .writeTimeslot(this.timeslot.startTime, this.timeslot.endTime)
      .subscribe(error => (this.error = error), id => (newId = id));
    return newId
  }
  
  onSelect(): void {
    if (this.timeFormat == "12hour") {
      this.twelveHourIsChecked = true;
    } else {
      this.twelveHourIsChecked = false;
    }
  }

  onSubmit(): void {
    // format timeslots
    var fullStart = this.eventDate
    var fullEnd = this.eventDate

    if (!this.twelveHourIsChecked) {
      var fullStart = this.format24HourTime(this.startHour, this.startMin, this.seconds)
      var fullEnd = this.format24HourTime(this.endHour, this.endHour, this.seconds)
    } else {
      var fullStart = this.format12HourTime(this.timeslot.startTime, this.seconds)
      var fullEnd = this.format12HourTime(this.timeslot.endTime, this.seconds)
    }

    this.timeslot.startTime = fullStart;
    this.timeslot.endTime = fullEnd;

    if (this.timeslot.startTime == "" || this.timeslot.endTime == "") {
      alert("Please enter a date and time for both fields");
      this.timeslotForm.reset();
    }

    // create new timeslot with user input
    var newTimeslot = new Timeslot(
      this.timeslot.startTime,
      this.timeslot.endTime
    );

    console.log(this.timeslot.startTime)
    console.log(this.timeslot.endTime)
    // pass new timeslot to timeslotService to send to database
    newTimeslot.id =  this.writeTimeslot(this.timeslot.startTime, this.timeslot.endTime)
    this.timeslotForm.reset();
    this.timeslots.push(newTimeslot);
  }

  deleteTimeslot(id): void {
    this.timeslotService
      .deleteTimeslot(id)
      .subscribe(error => (this.error = error));
    console.log("The following Timeslot Deleted :", this.timeslotForm.value);
    this.timeslots = this.timeslots.filter(item => item.id !== id);
  }

  updateTimeslot(): void {
    this.timeslot.isEditable = false;
    if (!this.twelveHourIsChecked) {
      var fullStart = this.format24HourTime(this.startHour, this.startMin, this.seconds)
      var fullEnd = this.format24HourTime(this.endHour, this.endMin, this.seconds)
    } else {
      var fullStart = this.format12HourTime(this.timeslot.startTime, this.seconds)
      var fullEnd = this.format12HourTime(this.timeslot.endTime, this.seconds)
    }

    this.timeslot.startTime = fullStart;
    this.timeslot.endTime = fullEnd;

    if (
      this.timeslot.startTime == "" ||
      this.timeslot.endTime == ""
    ) {
      alert("Please enter a date and time for both fields");
      this.timeslotForm.reset();
    }

    this.timeslotService
      .updateTimeslot(this.timeslot)
      .subscribe(
        error => (this.error = error),
        id => (this.timeslot.id = id)
      );

    this.getAllTimeslots();
    window.location.reload();
  }

  showEdit(timeslot: Timeslot): void {
    timeslot.isEditable = true;
    console.log(timeslot)
    this.timeslot.id = timeslot.id;
  }

  cancel(timeslot: Timeslot): void {
    timeslot.isEditable = false;
    this.timeslotForm.reset();
  }

  formatDate(timeslotValue: string): Date {
    var newTimeslotValue = timeslotValue.slice(0, -1);
    var newDate = new Date(newTimeslotValue);
    return newDate;
  }

  format24HourTime(hour: string, min: string, sec): string {
    return `${this.eventDate}T${hour}:${min}${sec}Z`
  }

  format12HourTime(startTime: string, sec: string): string {
    return `${this.eventDate}T${startTime}${sec}Z`
  }
}
