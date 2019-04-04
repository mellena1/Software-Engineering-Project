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

  login = new Login("", "");
  error: any;

  loginForm = new FormGroup({
    username: new FormControl(""),
    password: new FormControl("")
  });

  ngOnInit() {}

  submitLogin(username: string, password: string): void {
    var newUser = new Login(this.login.username, this.login.password);
    console.log("this is the user: " + this.login.username);
    this.loginService.writeLogin(this.login.password).subscribe(
      //response => {} not sure if I have a response here
      error => {
        this.error = error;
        console.log(this.error);
      }
    );
    this.loginService.setCookie(this.login.username);
    console.log("hello this is the cookie: " + this.loginService.getCookie()); //prints username
    //console.log("deleting the cookie: " + this.loginService.deleteCookie());
    this.loginForm.reset();
  }
}
