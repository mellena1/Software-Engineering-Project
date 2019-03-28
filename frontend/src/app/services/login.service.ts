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
import { WriteResponse } from "./writeResponse";

@Injectable({
  providedIn: "root"
})

export class LoginService {
	constructor(private http: HttpClient) {}
	private apiURL = environment.apiURL;
	jsonHeaders = new HttpHeaders().set("Content-Type", "application/json");

	
}