import { Component, Input, OnInit, ChangeDetectionStrategy } from "@angular/core";

import { ViewCell } from "ng2-smart-table";

import { TimeslotGlobals } from '../globals/timeslot.global';

import * as moment from "moment";

@Component({
  selector: 'timeslot',
  template: `
    <div *ngIf="timeslotGlobals.twelveHourIsChecked" style="font-size: 20px">
        {{ format12Hour(this.value) }}
    </div>
    <div *ngIf="!timeslotGlobals.twelveHourIsChecked" style="font-size: 20px">
        {{ format24Hour(this.value) }}
    </div>
  `
})
export class TimeslotRenderComponent implements ViewCell, OnInit {
  @Input() value: string;
  @Input() rowData: any;

  twelveHourIsChecked: boolean;

  constructor(private timeslotGlobals: TimeslotGlobals) {
    this.twelveHourIsChecked = timeslotGlobals.twelveHourIsChecked;
  }

  ngOnInit() { }

  format12Hour(time: string): string {
    return moment(time, 'YYYY-MM-DDTHH:mm:ss').format('h:mm a');
  }

  format24Hour(time: string): string {
    return moment(time, 'YYYY-MM-DDTHH:mm:ss').format('H:mm');
  }
}
