import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Timeslot } from '../data_models/timeslot'

@Injectable({
  providedIn: 'root'
})
export class TimeslotService {

  constructor(private http: HttpClient) { }

  getAllTimeslots() {

  }

  getTimeslot(id: number) {
    
  }
}
