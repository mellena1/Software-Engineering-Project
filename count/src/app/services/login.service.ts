import { Injectable } from "@angular/core";
import {
  HttpClient,
  HttpHeaders,
  HttpErrorResponse,
  HttpParams
} from "@angular/common/http";
import { environment } from "../../environments/environment";
import { Observable, throwError as observableThrowError } from "rxjs";
import { catchError, map } from "rxjs/operators";

import { Login } from "../data_models/login";
import { WriteResponse } from "./writeResponse"; //not sure if i can use this

@Injectable({
  providedIn: "root"
})
export class LoginService {
  constructor(private http: HttpClient) {}
  private apiUrl = environment.apiUrl;
  jsonHeaders = new HttpHeaders().set("Content-Type", "application/json");

  writeLogin(password: string) {
    var obj = { password: password };
    console.log(obj);
    return this.http.post<WriteResponse>(this.apiUrl + "/login", obj);
  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || "Server error");
  }

  setCookie(username: string) {
    document.cookie = "username = " + username;
  }

  getCookie() {
    return document.cookie;
  }

  deleteCookie() {
    const date = new Date();

    // Set it expire in -1 days
    date.setTime(date.getTime() + -1 * 24 * 60 * 60 * 1000);

    // Set it
    document.cookie = "username=; expires=" + date.toUTCString();
  }
}
