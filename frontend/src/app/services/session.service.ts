import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { map } from 'rxjs/operators'

import { Session } from '../data_models/session'

@Injectable({
  providedIn: 'root'
})
export class SessionService {
  private apiUrl = environment.apiUrl;
  
  constructor(private http: HttpClient) { }

  getAllSessions() {

  }

  getSession(id: number) {

  }

  writeSession() {

  }

  updateSession() {

  }

  deleteSession() {

  }

}
