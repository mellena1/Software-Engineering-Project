import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse, HttpParams} from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, throwError as observableThrowError } from 'rxjs';
import { catchError, map } from 'rxjs/operators'

import { Timeslot } from '../data_models/timeslot'
import { headersToString } from 'selenium-webdriver/http';
import { httpClientInMemBackendServiceFactory } from 'angular-in-memory-web-api';
import { stringify } from 'querystring';

@Injectable({
  providedIn: 'root'
})
export class TimeslotService {
  timeslot: Timeslot;

  private apiUrl = environment.apiUrl;
  
  constructor(private http: HttpClient) { }

  getAllTimeslots() {
    return this.http
    .get<Timeslot[]>(this.apiUrl + '/timeslot')
    .pipe(map(data => data), catchError(this.handleError));
  }

  getTimeslot(id: number) {
    let params = new HttpParams().set('id', stringify(id));
    return this.http
      .get<Timeslot>(this.apiUrl + '/timeslot', {params: params})
      .pipe(map(data => data), catchError(this.handleError));
  }

  writeTimeslot(timeslot: Timeslot) {
    return this.http.post<Timeslot>(this.apiUrl + '/timeslot', timeslot)
    .pipe(map(data => data), catchError(this.handleError));
  }

  updateTimeslot(updatedTimeslot: Timeslot) {
    return this.http.put<Timeslot>(this.apiUrl + '/timeslot', updatedTimeslot)
    .pipe(map(data => data), catchError(this.handleError));
  }

  deleteTimeslot(id: number) {
    let params = new HttpParams().set('id', stringify(id));
    return this.http.delete(this.apiUrl + '/timeslot', {params: params})
    .pipe(map(data => data), catchError(this.handleError));
  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || 'Server error');
  }
}
