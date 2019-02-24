import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Speaker } from '../data_models/speaker'

@Injectable({
  providedIn: 'root'
})
export class SpeakerService {

  constructor(private http: HttpClient) { }

  getAllSpeakers() {

  }

  getSpeaker(id: number) {
    
  }
}
