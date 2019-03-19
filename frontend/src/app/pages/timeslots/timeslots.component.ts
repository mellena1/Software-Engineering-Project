import { Component, OnInit } from "@angular/core";
import { Timeslot } from "src/app/data_models/timeslot";
import { TimeslotService } from "src/app/services/timeslot.service";
import { FormGroup, FormControl, Validators } from "@angular/forms";
import { format } from "url";

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
  currentTimeslot = new Timeslot("", "");
  date = new Date();
  currentDate: any;
  seconds: any;
  timeFormat: any;
  checked: any;

  startHour: any;
  startMin: any;
  endHour: any;
  endMin: any;

  currentStartHour: any;
  currentStartMin: any;
  currentEndHour: any;
  currentEndMin: any;

  timeslotForm = new FormGroup({
    timeslotStart: new FormControl(""),
    timeslotEnd: new FormControl("")
  });

  ngOnInit() {
    this.getAllTimeslots();
    this.currentDate = this.getCurrentDate();
    this.seconds = ":00";
    this.timeFormat = "12hour";
    this.checked = true;
    this.startHour = "00";
    this.startMin = "00";
    this.endHour = "00";
    this.endMin = "00";
    this.currentStartHour = "00";
    this.currentStartMin = "00";
    this.currentEndHour = "00";
    this.currentEndMin = "00";
  }

  getAllTimeslots(): void {
    this.timeslotService
      .getAllTimeslots()
      .subscribe(
        timeslots => (this.timeslots = timeslots),
        error => (this.error = error)
      );
  }

  onSelect(): void{
    if(this.timeFormat == "12hour"){
      this.checked = true;
    }
    else{
      this.checked = false;
    }
  }

  onSubmit(): void {

    var fullStart = this.currentDate.concat(" ");
    var fullEnd = this.currentDate.concat(" ");

    if(!this.checked) {

      fullStart = fullStart.concat(this.startHour).concat(":").concat(this.startMin).concat(this.seconds);
      fullEnd = fullEnd.concat(this.endHour).concat(":").concat(this.endMin).concat(this.seconds);
        
    }
    else {

      fullStart = fullStart.concat(this.timeslot.startTime).concat(this.seconds);
      fullEnd = fullEnd.concat(this.timeslot.endTime).concat(this.seconds);

    }
    
    this.timeslot.startTime = fullStart;
    this.timeslot.endTime = fullEnd;

    if (this.timeslot.startTime == "" || this.timeslot.endTime == "") {
      alert("Please enter a date and time for both fields");
      this.timeslotForm.reset();
    }

    console.log(this.timeslot.startTime);

    var newTimeslot = new Timeslot(
      this.timeslot.startTime,
      this.timeslot.endTime
    );

    this.timeslotService
      .writeTimeslot(this.timeslot.startTime, this.timeslot.endTime)
      .subscribe(error => (this.error = error), id => (newTimeslot.id = id));
    console.log("Timeslot Submitted!", this.timeslotForm.value);
    this.timeslotForm.reset();

    this.timeslots.push(newTimeslot);
  }

  deleteTimeslot(timeslotid): void {
    if (confirm("Are you sure you want to remove it?")) {
      this.timeslotService
        .deleteTimeslot(timeslotid)
        .subscribe(error => (this.error = error));
      console.log("The following Timeslot Deleted :", this.timeslotForm.value);
      this.timeslots = this.timeslots.filter(item => item.id !== timeslotid);
    }
  }

  updateTimeslot(timeslot: Timeslot): void {
    this.currentTimeslot.isEditable = false;

    var fullStart = this.currentDate.concat(" ");
    var fullEnd = this.currentDate.concat(" ");
    if(!this.checked) {

      fullStart = fullStart.concat(this.currentStartHour).concat(":").concat(this.currentStartMin).concat(this.seconds);
      fullEnd = fullEnd.concat(this.currentEndHour).concat(":").concat(this.currentEndMin).concat(this.seconds);
      
    }
    else {

      fullStart = fullStart.concat(this.timeslot.startTime).concat(this.seconds);
      fullEnd = fullEnd.concat(this.timeslot.endTime).concat(this.seconds);

    }

    this.currentTimeslot.startTime = fullStart;
    this.currentTimeslot.endTime = fullEnd;

    if (this.currentTimeslot.startTime == "" || this.currentTimeslot.endTime == "") {
      alert("Please enter a date and time for both fields");
      this.timeslotForm.reset();
    }

    this.timeslotService
       .updateTimeslot(this.currentTimeslot)
       .subscribe(
          error => (this.error = error),
          id => (this.currentTimeslot.id = id)
        );
    console.log("The following Timeslot Udpated :", this.timeslotForm.value);
    //this.timeslots = this.timeslots.filter(item => item.id !== timeslotid);
    this.timeslotForm.reset();
  }

  showEdit(timeslot: Timeslot): void {
    timeslot.isEditable = true;
    this.currentTimeslot.id = timeslot.id;
  }

  cancel(timeslot: Timeslot): void {
    timeslot.isEditable = false;
    this.timeslotForm.reset();
  }

  getCurrentDate(): String{
    var year = this.date.getUTCFullYear().toString();
    var day = this.date.getDate().toString();
    var m = this.date.getUTCMonth() + 1;
    var month = m.toString();

    if(Number(day) < 10){
      day = "0".concat(day);
    }
    if(Number(month) < 10){
      month = "0".concat(month);
    }

    return year.concat("-").concat(month).concat("-").concat(day);
  }
}
