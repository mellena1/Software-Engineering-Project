import { Injectable } from "@angular/core";
import {
  HttpClient,
  HttpHeaders,
  HttpParams
} from "@angular/common/http";
import { environment } from "../../environments/environment";
import { Observable } from "rxjs";

import { Timeslot } from "../data_models/timeslot";
import { WriteResponse } from "./writeResponse";

@Injectable({
  providedIn: "root"
})
export class TimeslotService {
  constructor(private http: HttpClient) {}
  private apiUrl = environment.apiUrl;
  jsonHeaders = new HttpHeaders().set("Content-Type", "application/json");

  getAllTimeslots(): Observable<Timeslot[]> {
    return this.http.get<Timeslot[]>(this.apiUrl + "/timeslots");
  }

  getTimeslot(id: number): Observable<Timeslot> {
    var params = new HttpParams().set("id", id.toString());
    return this.http.get<Timeslot>(this.apiUrl + "/timeslot", {
      params: params
    });
  }

  writeTimeslot(startTime: string, endTime: string): Observable<WriteResponse> {
    var obj = { startTime: startTime, endTime: endTime };
    return this.http.post<WriteResponse>(this.apiUrl + "/timeslot", obj);
  }

  updateTimeslot(updatedTimeslot: Timeslot): Observable<any> {
    var obj = {
      id: updatedTimeslot.id,
      startTime: updatedTimeslot.startTime,
      endTime: updatedTimeslot.endTime
    };
    return this.http.put(this.apiUrl + "/timeslot", obj);
  }

  deleteTimeslot(id: number): Observable<any> {
    var params = new HttpParams().set("id", id.toString());
    return this.http.delete(this.apiUrl + "/timeslot", {
      params: params
    });
  }
}
