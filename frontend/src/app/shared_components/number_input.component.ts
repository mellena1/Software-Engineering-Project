import { Component, ViewChild, ElementRef, AfterViewInit } from "@angular/core";

import { Cell, DefaultEditor, Editor } from "ng2-smart-table";

@Component({
  template: `
    <input
      type="number"
      [(ngModel)]="stringNumber"
      style="font-size: 20px;"
    />
    <div [hidden]="true" [innerHTML]="parseToInt()" #htmlValue></div>
  `
})
export class NumberInputComponent extends DefaultEditor {
  stringNumber: string;

  ngOnInit() {
    this.stringNumber = this.cell.getValue();
  }

  parseToInt() {
    this.cell.newValue = parseInt(this.stringNumber);
  }
}
