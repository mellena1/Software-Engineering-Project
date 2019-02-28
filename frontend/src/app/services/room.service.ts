import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse, HttpParams } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, throwError as observableThrowError } from 'rxjs';
import { catchError, map } from 'rxjs/operators'

import { Room } from '../data_models/room'
import { stringify } from '@angular/core/src/render3/util';
import { headersToString } from 'selenium-webdriver/http';

@Injectable({
  providedIn: 'root'
})
export class RoomService {
  private apiUrl = environment.apiUrl;
  
  

  constructor(private http: HttpClient) { }
  headers = new HttpHeaders();

  
  


  getAllRooms() {
    return this.http
      .get<Room[]>(this.apiUrl + '/rooms')
      .pipe(map(data => data), catchError(this.handleError));
  }

  getARoom(id: number) :Observable<Room>{
    const data = new FormData();
    this.headers.append('id', 'number');
    this.headers.set('id', stringify(id));
    data.append('id', stringify(id));

    return this.http
    .get<Room>(this.apiUrl + '/room', data)
    .pipe(map(data=>data), catchError(this.handleError));
    
  }

  writeRoom(room: Room) {
    return this.http.get(this.apiUrl + '/room')
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
