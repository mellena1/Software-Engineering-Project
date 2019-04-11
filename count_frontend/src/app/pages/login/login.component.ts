import { Component, OnInit } from '@angular/core';
import { LoginService } from 'src/app/services/login.service';
import { User } from 'src/app/data_models/user';
import { Router } from "@angular/router";


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(private loginService: LoginService, private router: Router) { }

  model = new User("", "");
  showWarning: boolean;

  ngOnInit() {
    this.model = new User("", "");
  }

  onSubmit() {
    if (this.loginService.login(this.model)) {
      this.router.navigate(["/count"])
    } else {
      this.showWarning = true;
    }
  }
}
