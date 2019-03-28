import { Component, OnInit } from "@angular/core";
import { FormControl, FormGroup, Validators } from "@angular/forms";
import { Login } from "src/app/data_models/login";
import { LoginService } from "src/app/services/login.service";

@Component({
	selector: "",
	templateUrl: "./login.component.html",
	styleUrls: ["./login.component.css"]
})


export class LoginComponent implements OnInit {
	constructor(private loginService: LoginService) {}


	submitLogin(): void {
		var newUser = new Login(
			this.login.username,
			this.login.password
		);
		this.loginService
		 .submitLogin(
		 	this.login.username,
		 	this.login.password
		 )
		 .subscribe(
		 	//response => {} not sure if I have a response here
		 	error => {
		 	  this.error = error;
		 	  console.log(this.error);
		 	}
		 );
	}
}
