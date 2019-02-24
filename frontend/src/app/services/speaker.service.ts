import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { map } from 'rxjs/operators'

import { Speaker } from '../data_models/speaker'

@Injectable({
  providedIn: 'root'
})
export class SpeakerService {
  private apiUrl = environment.apiUrl;
  
  constructor(private http: HttpClient) { }

  getAllSpeakers() {

  }

  getSpeaker(id: number) {
    
  }
}
