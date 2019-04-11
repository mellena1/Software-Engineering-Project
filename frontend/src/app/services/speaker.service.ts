import { Injectable } from "@angular/core";
import {
  HttpClient,
  HttpHeaders,
  HttpParams
} from "@angular/common/http";
import { environment } from "../../environments/environment";
import { Observable } from "rxjs";

import { Speaker } from "../data_models/speaker";
import { WriteResponse } from "./writeResponse";

@Injectable({
  providedIn: "root"
})
export class SpeakerService {
  constructor(private http: HttpClient) {}
  private apiUrl = environment.apiUrl;
  jsonHeaders = new HttpHeaders().set("Content-Type", "application/json");

  getAllSpeakers(): Observable<Speaker[]> {
    return this.http.get<Speaker[]>(this.apiUrl + "/speakers");
  }

  getSpeaker(id: number): Observable<Speaker> {
    var params = new HttpParams().set("id", id.toString());
    return this.http.get<Speaker>(this.apiUrl + "/speaker", {
      params: params
    });
  }

  writeSpeaker(
    firstName: string,
    lastName: string,
    email: string
  ): Observable<WriteResponse> {
    var obj = { firstName: firstName, lastName: lastName, email: email };
    return this.http.post<WriteResponse>(this.apiUrl + "/speaker", obj);
  }

  updateSpeaker(updatedSpeaker: Speaker): Observable<any> {
    return this.http.put(this.apiUrl + "/speaker", updatedSpeaker);
  }

  deleteSpeaker(id: number): Observable<any> {
    var params = new HttpParams().set("id", id.toString());
    return this.http.delete(this.apiUrl + "/speaker", {
      params: params
    });
  }
}
