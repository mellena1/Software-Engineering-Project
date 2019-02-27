import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse} from '@angular/common/http';
import { environment } from '../../environments/environment';
//import { map } from 'rxjs/operators';
import { Observable, throwError as observableThrowError } from 'rxjs';
import { catchError, map } from 'rxjs/operators'

import { Timeslot } from '../data_models/timeslot'

@Injectable({
  providedIn: 'root'
})
export class TimeslotService {
  private apiUrl = environment.apiUrl;
  
  constructor(private http: HttpClient) { }

  getAllTimeslots() {
    return this.http
    .get<Timeslot[]>(this.apiUrl + '/timeslot')
    .pipe(map(data => data), catchError(this.handleError));
  }

  getTimeslot(id: number) {
    return this.http
    .get<Timeslot>(this.apiUrl + '/timeslot')
  }

  writeTimeslot() {

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
