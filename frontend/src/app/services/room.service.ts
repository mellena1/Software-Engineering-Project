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

import { Room } from "../data_models/room";

@Injectable({
  providedIn: "root"
})
export class RoomService {
  constructor(private http: HttpClient) {}
  private apiUrl = environment.apiUrl;
  jsonHeaders = new HttpHeaders().set("Content-Type", "application/json");

  getAllRooms() {
    return this.http.get<Room[]>(this.apiUrl + "/rooms").pipe(
      map(data => data),
      catchError(this.handleError)
    );
  }

  getARoom(id: number) {
    var params = new HttpParams().set("id", id.toString());
    return this.http.get<Room>(this.apiUrl + "/room", {
      params: params
    });
  }

  writeRoom(name: string, capacity: number) {
    var obj = { name: name, capacity: capacity };
    return this.http.post(this.apiUrl + "/room", obj);
  }

  updateRoom(updatedRoom: Room) {
    return this.http.post(this.apiUrl + "/room", {
      headers: this.jsonHeaders,
      body: JSON.stringify(updatedRoom)
    });
  }

  deleteRoom(id: number) {
    var params = new HttpParams().set("id", id.toString());
    return this.http.delete(this.apiUrl + "/room", {
      params: params
    });
  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || "Server error");
  }
}
