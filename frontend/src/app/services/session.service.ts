import { Injectable } from "@angular/core";
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http";
import { environment } from "../../environments/environment";
import { Observable } from "rxjs";

import { Session } from "../data_models/session";
import { WriteResponse } from "./writeResponse";

@Injectable({
  providedIn: "root"
})
export class SessionService {
  constructor(private http: HttpClient) {}
  private apiUrl = environment.apiUrl;
  jsonHeaders = new HttpHeaders().set("Content-Type", "application/json");

  getAllSessions(): Observable<Session[]> {
    return this.http.get<Session[]>(this.apiUrl + "/sessions");
  }

  getSession(id: number): Observable<Session> {
    var params = new HttpParams().set("id", id.toString());
    return this.http.get<Session>(this.apiUrl + "/session", {
      params: params
    });
  }

  writeSession(
    name: string,
    roomID: number,
    speakerID: number,
    timeslotID: number
  ): Observable<WriteResponse> {
    var obj = {
      sessionName: name,
      roomID: roomID,
      speakerID: speakerID,
      timeslotID: timeslotID
    };
    return this.http.post<WriteResponse>(this.apiUrl + "/session", obj);
  }

  updateSession(updatedSession: Session): Observable<any> {
    var obj = {
      roomID: updatedSession.room.id,
      sessionID: updatedSession.id,
      sessionName: updatedSession.name,
      speakerID: updatedSession.speaker.id,
      timeslotID: updatedSession.timeslot.id
    };
    return this.http.put(this.apiUrl + "/session", obj);
  }

  deleteSession(id: number): Observable<any> {
    var params = new HttpParams().set("id", id.toString());
    return this.http.delete(this.apiUrl + "/session", {
      params: params
    });
  }
}
