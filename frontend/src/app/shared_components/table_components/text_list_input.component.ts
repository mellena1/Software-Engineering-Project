import { Component, OnInit } from "@angular/core";

import { DefaultEditor } from "ng2-smart-table";

@Component({
  template: `
    <select
      [ngClass]="inputClass"
      class="form-control"
      [(ngModel)]="cell.newValue"
      [name]="cell.getId()"
      [disabled]="!cell.isEditable()"
      (click)="onClick.emit($event)"
      (keydown.enter)="onEdited.emit($event)"
      (keydown.esc)="onStopEditing.emit()"
    >
      <option
        *ngFor="let option of cell.getColumn().getConfig()?.list"
        [ngValue]="option.value"
        [selected]="option.value === cell.getValue()"
        >{{ option.title }}
      </option>
    </select>
  `
})
export class TextListInputComponent extends DefaultEditor implements OnInit {
  constructor() {
    super();
  }

  ngOnInit() {}
}
