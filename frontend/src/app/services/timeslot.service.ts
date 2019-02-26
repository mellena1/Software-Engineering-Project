import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { map } from 'rxjs/operators'

import { Timeslot } from '../data_models/timeslot'

@Injectable({
  providedIn: 'root'
})
export class TimeslotService {
  private apiUrl = environment.apiUrl;
  
  constructor(private http: HttpClient) { }

  getAllTimeslots() {

  }

  getTimeslot(id: number) {
    
  }

  writeTimeslot() {

  }

  updateTimeslot() {

  }

  deleteTimeslot() {

  }
}
