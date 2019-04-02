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
  private apiURL = environment.apiURL;
  jsonHeaders = new HttpHeaders().set("Content-Type", "application/json");

  submitLogin(username: string, password: string) {
    var obj = { username: username, password: password };
    return this.http.post<WriteResponse>(this.apiUrl + "/login", obj); //need to double check path
  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || "Server error");
  }
}
