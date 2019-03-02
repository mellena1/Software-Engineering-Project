import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, throwError as observableThrowError } from 'rxjs';
import { catchError, map } from 'rxjs/operators'

import { Timeslot } from '../data_models/timeslot'

@Injectable({
  providedIn: 'root'
})
export class TimeslotService {
  constructor(private http: HttpClient) { }
  private apiUrl = environment.apiUrl;
  
  getAllTimeslots() {
    return this.http
      .get<Timeslot[]>(this.apiUrl + '/timeslots')
      .pipe(map(data => data), catchError(this.handleError))
  }

  getTimeslot(id: number) {

  }

  writeTimeslot(timeslot: Timeslot) {

  }

  updateTimeslot(updatedTimeslot: Timeslot) {

  }

  deleteTimeslot(id: number) {

  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || 'Server error');
  }
}
