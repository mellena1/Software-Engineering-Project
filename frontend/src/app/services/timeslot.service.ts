import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse} from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, throwError as observableThrowError } from 'rxjs';
import { catchError, map } from 'rxjs/operators'

import { Timeslot } from '../data_models/timeslot'
import { headersToString } from 'selenium-webdriver/http';
import { httpClientInMemBackendServiceFactory } from 'angular-in-memory-web-api';

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
 
  }

  writeTimeslot(timeslot: Timeslot) {
    return this.http.post(this.apiUrl + '/timeslot', timeslot)
  }

  updateTimeslot() {

  }

  deleteTimeslot() {

  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || 'Server error');
  }
}
