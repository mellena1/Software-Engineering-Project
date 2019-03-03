import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, throwError as observableThrowError } from 'rxjs';
import { catchError, map } from 'rxjs/operators'

import { Room } from '../data_models/room'

@Injectable({
  providedIn: 'root'
})
export class RoomService {
  private apiUrl = environment.apiUrl;

  constructor(private http: HttpClient) { }

  getAllRooms() {
    return this.http
      .get<Room[]>(this.apiUrl + '/room')
      .pipe(map(data => data), catchError(this.handleError));
  }

  getARoom(id: number) {
    
  }

  writeRoom() {

  }

  updateRoom() {

  }

  deleteRoom() {

  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || 'Server error');
  }
}
