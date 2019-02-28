import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, throwError as observableThrowError } from 'rxjs';
import { catchError, map } from 'rxjs/operators'

import { Speaker } from '../data_models/speaker'

@Injectable({
  providedIn: 'root'
})

export class SpeakerService {
  private apiUrl = environment.apiUrl;
  
  constructor(private http: HttpClient) { }

  getAllSpeakers() {
    return this.http
      .get<Speaker[]>(this.apiUrl + '/speakers')
      .pipe(map(data => data), catchError(this.handleError))
  }
  
  getSpeaker(id: number) {
    
  }

  writeSpeaker(speaker: Speaker) {

  }

  updateSpeaker(updateSpeaker: Speaker) {

  }

  deleteSpeaker(id: number) {

  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || 'Server error');
  }

}
