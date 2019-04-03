import { Component, OnInit } from "@angular/core";
import { FormControl, FormGroup, Validators } from "@angular/forms";
import { Count } from "src/app/data_models/count";
import { CountService } from "src/app/services/count.service";

@Component({
  selector: "",
  templateUrl: "./counter.component.html",
  styleUrls: ["./counter.component.css"]
})
export class CounterComponent implements OnInit {
  constructor(private countService: CountService) {}
  // username: document.cookie; //need to set username here from login
  count: Count;
  error: any;

  countForm = new FormGroup({
    userName: new FormControl(""), //document.cookie, //passed in
    count: new FormControl(""),
    time: new FormControl(""), //going to be button, not form
    sessionID: new FormControl("") //this will be passed as well
  });

  ngOnInit() { }

  getACount(sessionID): void {
    this.countService
      .getACount(sessionID)
      .subscribe(count => (this.count = count), error => (this.error = error));
  }

  writeACount(): void {
    var newCount = new Count(
      this.count.count,
      this.count.time,
      this.count.userName
    );
    this.countService
      .writeACount(newCount)
      .subscribe(
        response => {
          this.countForm.reset();
          console.log("Count Submitted");
        },
        error => {
          this.error = error;
          console.log(this.error);
        }
      );
  }

  updateCount(): void {
    //not sure where to take this
    // var index = this.count.findIndex(item => item.id);
  }
}
