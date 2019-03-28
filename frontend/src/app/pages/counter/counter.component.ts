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
	username: //need to set username here from login
	error: any;

	countForm = new FormGroup({
		userName: //passed in
		count: new FormControl(""),
		time: new //going to be button, not form
		sessionID: new //this will be passed as well
	});

	getACount(sessionID): void {
	  this.CountService()
	  .getACount(sessionID)
	  .subscribe(
	  	count => (this.count = count)
	  	error => (this.error = error)
	  	);
	}

	writeACount(): void {
	  var newCount = new Count(
	  	this.count.count,
	  	this.count.username,
	  	this.count.sessionID,
	  	this.count.time
	  );
	  this.countService
	  	.writeACount(
	  	  this.count.count,
	  	  this.count.username,
	  	  this.count.sessionID,
	  	  this.count.time
	  	)
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

	updateCount(): void { //not sure where to take this
		var index = this.count.findIndex(
			item => item.id
		);


	}


}
