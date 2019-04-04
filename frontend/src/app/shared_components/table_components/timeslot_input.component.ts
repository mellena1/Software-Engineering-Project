import { Component, OnInit } from "@angular/core";

import { TimeslotGlobals } from '../../globals/timeslot.global';

import { DefaultEditor } from "ng2-smart-table";

import * as moment from "moment";

@Component({
  template: `
    <ngb-timepicker *ngIf="timeslotGlobals.twelveHourIsChecked" [(ngModel)]="time" [meridian]=true></ngb-timepicker>
    <ngb-timepicker *ngIf="!timeslotGlobals.twelveHourIsChecked" [(ngModel)]="time" [meridian]=false></ngb-timepicker>
    <div [hidden]="true" [innerHTML]="formatToRFC3331()" #htmlValue></div>
  `,
  styleUrls: ["./table.components.css"]
})
export class TimeslotInputComponent extends DefaultEditor implements OnInit {
  time: object;

  constructor(private timeslotGlobals: TimeslotGlobals) {
    super();
  }

  ngOnInit() { 
    this.time = this.RFC3331ToTimeInput(this.cell.getValue());
  }

  formatToRFC3331(): void {
    this.cell.newValue = moment(this.time).format('YYYY-MM-DDTHH:mm:ss') + 'Z';
  }

  RFC3331ToTimeInput(time: string): object {
    var obj = moment(time, 'YYYY-MM-DDTHH:mm:ss').toObject();
    return {hour: obj.hours, minute: obj.minutes};
  }
}
