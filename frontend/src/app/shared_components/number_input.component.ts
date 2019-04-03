import { Component, ViewChild, ElementRef, AfterViewInit } from "@angular/core";

import { Cell, DefaultEditor, Editor } from "ng2-smart-table";

@Component({
  template: `
    <input
      type="number"
      [(ngModel)]="this.cell.newValue"
      style="font-size: 20px;"
    />
  `
})
export class NumberInputComponent extends DefaultEditor {
  ngOnInit() {
    var curVal = this.cell.getValue();
    this.cell.newValue = parseInt(curVal);
  }
}
