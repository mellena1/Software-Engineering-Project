import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse, HttpParams } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, throwError as observableThrowError } from 'rxjs';
import { catchError, map } from 'rxjs/operators'

import { Session } from '../data_models/session'

@Injectable({
  providedIn: 'root'
})
export class SessionService {
  constructor(private http: HttpClient) { }
  private apiUrl = environment.apiUrl;
  
  getAllSessions() {
    return this.http
      .get<Session[]>(this.apiUrl + '/sessions')
      .pipe(map(data => data), catchError(this.handleError));
  }

  getSession(id: number) {
    const params = new HttpParams()
      .set('id', id.toString());
    return this.http.get<Session>(this.apiUrl + '/session', {
      params: params
    });
  }

  writeSession(session: Session) {

  }

  updateSession(updatedSession: Session) {

  }

  deleteSession(id: number) {
    const params = new HttpParams()
      .set('id', id.toString());
    return this.http.delete<Session>(this.apiUrl + '/session', {
      params: params
    });
  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || 'Server error');
  }
}
