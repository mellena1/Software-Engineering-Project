import { Component, Input, OnInit } from "@angular/core";

import { ViewCell } from "ng2-smart-table";

import { TimeslotGlobals } from "../../globals/timeslot.global";

@Component({
  selector: "timeslot",
  template: `
    <div class="table-text">{{ formatTimeForCell(this.value) }}</div>
  `,
  styleUrls: ["./table.components.css"]
})
export class TimeslotRenderComponent implements ViewCell, OnInit {
  @Input() value: string;
  @Input() rowData: any;

  constructor(private timeslotGlobals: TimeslotGlobals) {}

  ngOnInit() {}

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

    if (TimeslotGlobals.isValidTime(startTime) && TimeslotGlobals.isValidTime(endTime)) {
      return `${this.formatTimeForCellSingleTime(
        startTime
      )}-${this.formatTimeForCellSingleTime(endTime)}`;
    } else {
      // will return whatever is valid, or '' if neither are valid
      return `${this.formatTimeForCellSingleTime(
        startTime
      )}${this.formatTimeForCellSingleTime(endTime)}`
    }
  }

  formatTimeForCellSingleTime(time: string): string {
    if (TimeslotGlobals.isValidTime(time)) {
      return `${this.timeslotGlobals.formatTime(time)}`;
    } else {
      return '';
    }
  }
}
