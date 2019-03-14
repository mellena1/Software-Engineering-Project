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

import { Speaker } from "../data_models/speaker";
import { WriteResponse } from "./writeResponse";

@Injectable({
  providedIn: "root"
})
export class SpeakerService {
  constructor(private http: HttpClient) {}
  private apiUrl = environment.apiUrl;
  jsonHeaders = new HttpHeaders().set("Content-Type", "application/json");

  getAllSpeakers() {
    return this.http.get<Speaker[]>(this.apiUrl + "/speakers").pipe(
      map(data => data),
      catchError(this.handleError)
    );
  }

  getSpeaker(id: number) {
    var params = new HttpParams().set("id", id.toString());
    return this.http.get<Speaker>(this.apiUrl + "/speaker", {
      params: params
    });
  }

  writeSpeaker(firstName: string, lastName: string, email: string) {
    var obj = { firstName: firstName, lastName: lastName, email: email };
    return this.http.post<WriteResponse>(this.apiUrl + "/speaker", obj);
  }

  updateSpeakers(updatedSpeaker: Speaker) {
    return this.http.post(this.apiUrl + "/speaker", updatedSpeaker)
  }

  deleteSpeaker(id: number) {
    var params = new HttpParams().set("id", id.toString());
    return this.http.delete(this.apiUrl + "/speaker", {
      params: params
    });
  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || "Server error");
  }
}
