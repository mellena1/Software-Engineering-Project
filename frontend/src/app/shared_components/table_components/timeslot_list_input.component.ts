import { Component, OnInit } from "@angular/core";

import { TimeslotGlobals } from "../../globals/timeslot.global";

import { DefaultEditor } from "ng2-smart-table";

import { TimeslotRenderHelpers } from "./timeslot_render.helpers";

@Component({
  template: `
  <select [ngClass]="inputClass"
    class="form-control"
    [(ngModel)]="cell.newValue"
    [name]="cell.getId()"
    [disabled]="!cell.isEditable()"
    (click)="onClick.emit($event)"
    (keydown.enter)="onEdited.emit($event)"
    (keydown.esc)="onStopEditing.emit()">
    <option *ngFor="let option of cell.getColumn().getConfig()?.list" [value]="option.value"
      [selected]="option.value === cell.getValue()">{{ formatTimeForCell(option.title) }}
    </option>
  </select>
  `
})
export class TimeslotListInputComponent extends DefaultEditor implements OnInit {
  constructor(private timeslotGlobals: TimeslotGlobals) {
    super();
  }

  ngOnInit() { }

  formatTimeForCell(time: string): string {
    return TimeslotRenderHelpers.formatTimeForCell(time, this.timeslotGlobals);
  }
}
