import { Component, Input, OnInit } from "@angular/core";

import { ViewCell } from "ng2-smart-table";

import { TimeslotGlobals } from "../../globals/timeslot.global";

import { TimeslotRenderHelpers } from "./timeslot_render.helpers";

@Component({
  selector: "timeslot",
  template: `
    <div class="table-text">
      {{ formatTimeForCell(this.value, this.timeslotGlobals) }}
    </div>
  `,
  styleUrls: ["./table.components.css"]
})
export class TimeslotRenderComponent implements ViewCell, OnInit {
  @Input() value: string;
  @Input() rowData: any;

  constructor(private timeslotGlobals: TimeslotGlobals) {}

  ngOnInit() {}

  formatTimeForCell(time: string): string {
    return TimeslotRenderHelpers.formatTimeForCell(time, this.timeslotGlobals);
  }
}
