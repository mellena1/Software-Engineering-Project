import { Injectable } from "@angular/core";
import {
  HttpClient,
  HttpParams
} from "@angular/common/http";
import { environment } from "../../environments/environment";
import { Observable } from "rxjs";

import { Room } from "../data_models/room";
import { WriteResponse } from "./writeResponse";

@Injectable({
  providedIn: "root"
})
export class RoomService {
  constructor(private http: HttpClient) {}
  private apiUrl = environment.apiUrl;

  getAllRooms(): Observable<Room[]> {
    return this.http
      .get<Room[]>(this.apiUrl + "/rooms");
  }

  getARoom(id: number): Observable<Room> {
    var params = new HttpParams().set("id", id.toString());
    return this.http.get<Room>(this.apiUrl + "/room", {
      params: params
    });
  }

  writeRoom(name: string, capacity: number): Observable<WriteResponse> {
    var obj = { name: name, capacity: capacity };
    return this.http.post<WriteResponse>(this.apiUrl + "/room", obj);
  }

  updateRoom(updatedRoom: Room): Observable<any> {
    return this.http.put(this.apiUrl + "/room", updatedRoom);
  }

  deleteRoom(id: number): Observable<any> {
    var params = new HttpParams().set("id", id.toString());
    return this.http.delete(this.apiUrl + "/room", {
      params: params
    });
  }
}
