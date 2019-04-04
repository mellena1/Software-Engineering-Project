import { Component, Input, OnInit } from "@angular/core";

import { ViewCell } from "ng2-smart-table";

import { TimeslotGlobals } from '../../globals/timeslot.global';

import * as moment from "moment";

@Component({
  selector: 'timeslot',
  template: `
    <div class="table-text">{{ formatTimeForCell(this.value) }}</div>
  `,
  styleUrls: ["./table.components.css"]
})
export class TimeslotRenderComponent implements ViewCell, OnInit {
  @Input() value: string;
  @Input() rowData: any;

  constructor(private timeslotGlobals: TimeslotGlobals) { }

  ngOnInit() { }

  formatTimeForCell(time: string): string {
    if (time.includes(" ")) {
      return this.formatTimeForCellStartDashEnd(time);
    } else {
      return this.formatTimeForCellSingleTime(time);
    }
  }

  formatTimeForCellStartDashEnd(time: string): string {
    var splitValue = time.split(" ");
    var startTime = splitValue[0];
    var endTime = splitValue[1];
    if (this.timeslotGlobals.twelveHourIsChecked) {
      return `${this.format12Hour(startTime)}-${this.format12Hour(endTime)}`;
    } else {
      return `${this.format24Hour(startTime)}-${this.format24Hour(endTime)}`;
    }
  }

  formatTimeForCellSingleTime(time: string): string {
    if (this.timeslotGlobals.twelveHourIsChecked) {
      return `${this.format12Hour(time)}`;
    } else {
      return `${this.format24Hour(time)}`;
    }
  }

  format12Hour(time: string): string {
    return moment(time, 'YYYY-MM-DDTHH:mm:ss').format('h:mm a');
  }

  format24Hour(time: string): string {
    return moment(time, 'YYYY-MM-DDTHH:mm:ss').format('H:mm');
  }
}
