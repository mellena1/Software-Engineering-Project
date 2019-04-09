import { Component, OnInit, Input } from "@angular/core";

import { TimeslotGlobals } from "../globals/timeslot.global";

import { LocalDataSource } from "ng2-smart-table";

@Component({
  selector: "twelve-twentyfour-hour-radio",
  template: `
    <input
      type="radio"
      name="timeformat"
      [value]="true"
      (change)="tableDataSource.refresh()"
      [(ngModel)]="timeslotGlobals.twelveHourIsChecked"
    />
    12 Hour

    <input
      type="radio"
      name="timeformat"
      [value]="false"
      (change)="tableDataSource.refresh()"
      [(ngModel)]="timeslotGlobals.twelveHourIsChecked"
    />
    24 Hour
  `
})
export class TwelveTwentyfourHourRadioComponent implements OnInit {
  @Input() tableDataSource: LocalDataSource;

  constructor(private timeslotGlobals: TimeslotGlobals) {}

  ngOnInit() {}
}
