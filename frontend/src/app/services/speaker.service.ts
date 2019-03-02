import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse, HttpParams } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable, throwError as observableThrowError } from 'rxjs';
import { catchError, map } from 'rxjs/operators'

import { Speaker } from '../data_models/speaker'

@Injectable({
  providedIn: 'root'
})

export class SpeakerService {
  constructor(private http: HttpClient) { }
  private apiUrl = environment.apiUrl;
  headers = new HttpHeaders().set('Content-Type', 'application/json')
  
  getAllSpeakers() {
    return this.http
      .get<Speaker[]>(this.apiUrl + '/speakers')
      .pipe(map(data => data), catchError(this.handleError))
  }
  
  getSpeaker(id: number) {
    const params = new HttpParams()
      .set('personId', id.toString());
    return this.http.delete(this.apiUrl + '/speaker', {
      headers: this.headers,
      params: params
      }
    );
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
