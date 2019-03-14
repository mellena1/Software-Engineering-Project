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

import { Timeslot } from "../data_models/timeslot";
import { WriteResponse } from "./writeResponse";

@Injectable({
  providedIn: "root"
})
export class TimeslotService {
  constructor(private http: HttpClient) {}
  private apiUrl = environment.apiUrl;
  jsonHeaders = new HttpHeaders().set("Content-Type", "application/json");

  getAllTimeslots() {
    return this.http.get<Timeslot[]>(this.apiUrl + "/timeslots").pipe(
      map(data => data),
      catchError(this.handleError)
    );
  }

  getTimeslot(id: number) {
    var params = new HttpParams().set("id", id.toString());
    return this.http.get<Timeslot>(this.apiUrl + "/timeslot", {
      params: params
    });
  }

  writeTimeslot(startTime: string, endTime: string) {
    var obj = { startTime: startTime, endTime: endTime };
    return this.http.post<WriteResponse>(this.apiUrl + "/timeslot", obj)
  }

  updateTimeslot(updatedTimeslot: Timeslot) {
    return this.http.post(this.apiUrl + "/timeslot", updatedTimeslot)
  }

  deleteTimeslot(id: number) {
    var params = new HttpParams().set("id", id.toString());
    return this.http.delete(this.apiUrl + "/timeslot", {
      params: params
    });
  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || "Server error");
  }
}
