import { Injectable } from '@angular/core';
import { environment } from "../../environments/environment";
import { User } from '../data_models/user';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor() { }
  private password = environment.password;

  private currentUser = new User("", "");

  login(user: User): boolean {
    if (this.password == user.password) {
      this.currentUser = user
      return true
    } else {
      return false
    }
  }

  getCurrentUsername(): string {
    return this.currentUser.username
  }

}
