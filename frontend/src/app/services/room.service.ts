import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { map } from 'rxjs/operators'

import { Room } from '../data_models/room'

@Injectable({
  providedIn: 'root'
})
export class RoomService {
  private apiUrl = environment.apiUrl;

  constructor(private http: HttpClient) { }

  getAllRooms() {
    return this.http
      .get<Room[]>(this.apiUrl + '/api/v1/getAllRooms')
      .pipe(map(data => data));
  }

  getRoom(id: number) {
    
  }
}
