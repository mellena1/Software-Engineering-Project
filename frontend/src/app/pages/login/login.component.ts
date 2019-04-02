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
    var newUser = new Login(this.login.username, this.login.password);
    this.loginService
      .submitLogin(this.login.username, this.login.password)
      .subscribe(
        //response => {} not sure if I have a response here
        error => {
          this.error = error;
          console.log(this.error);
        }
      );
  }
}

export function setCookie(username: string, val: string): void {
  document.cookie = username + "=" + val;
}

export function getCookie(username: string) {
  return document.cookie;
}

export function deleteCookie(username: string) {
  const date = new Date();

  // Set it expire in -1 days
  date.setTime(date.getTime() + -1 * 24 * 60 * 60 * 1000);

  // Set it
  document.cookie = name + "=; expires=" + date.toUTCString();
}
