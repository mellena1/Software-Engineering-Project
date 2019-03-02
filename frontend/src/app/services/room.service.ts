import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse, HttpParams } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, throwError as observableThrowError } from 'rxjs';
import { catchError, map } from 'rxjs/operators'

import { Room } from '../data_models/room'

@Injectable({
  providedIn: 'root'
})
export class RoomService {
  constructor(private http: HttpClient) { }
  private apiUrl = environment.apiUrl;
  

  getAllRooms() {
    return this.http
      .get<Room[]>(this.apiUrl + '/rooms')
      .pipe(map(data => data), catchError(this.handleError));
  }

  getARoom(id: number) {

  }

  writeRoom(room: Room) {
  
  }

  updateRoom(updatedRoom: Room) {
    
  }

  deleteRoom(id: number) {

  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || 'Server error');
  }
}