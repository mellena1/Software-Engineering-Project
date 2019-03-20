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

import { Session } from "../data_models/session";
import { WriteResponse } from "./writeResponse";

@Injectable({
  providedIn: "root"
})
export class SessionService {
  constructor(private http: HttpClient) {}
  private apiUrl = environment.apiUrl;
  jsonHeaders = new HttpHeaders().set("Content-Type", "application/json");

  getAllSessions() {
    return this.http.get<Session[]>(this.apiUrl + "/sessions").pipe(
      map(data => data),
      catchError(this.handleError)
    );
  }

  getSession(id: number) {
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
  ) {
    var obj = {
      sessionName: name,
      roomID: roomID,
      speakerID: speakerID,
      timeslotID: timeslotID
    };
    return this.http.post<WriteResponse>(this.apiUrl + "/session", obj);
  }

  updateSession(updatedSession: Session) {
    var obj = {
      roomID: updatedSession.room.id,
      sessionID: updatedSession.id,
      sessionName: updatedSession.name,
      speakerID: updatedSession.speaker.id,
      timeslotID: updatedSession.timeslot.id
    }
    return this.http.put(this.apiUrl + "/session", obj);
  }

  deleteSession(id: number) {
    var params = new HttpParams().set("id", id.toString());
    return this.http.delete(this.apiUrl + "/session", {
      params: params
    });
  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || "Server error");
  }
}
