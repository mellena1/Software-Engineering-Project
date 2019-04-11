import { Component, OnInit } from "@angular/core";

import { DefaultEditor } from "ng2-smart-table";

@Component({
  template: `
    <input
      type="number"
      [(ngModel)]="stringNumber"
      (change)="parseToInt()"
      class="table-text"
    />
  `,
  styleUrls: ["./table.components.css"]
})
export class NumberInputComponent extends DefaultEditor implements OnInit {
  stringNumber: string;

  ngOnInit() {
    this.stringNumber = this.cell.getValue();
    this.parseToInt();
  }

  parseToInt() {
    var asInt = parseInt(this.stringNumber);
    if (Number.isNaN(asInt)) {
      this.cell.newValue = null;
    } else {
      this.cell.newValue = asInt;
    }
  }
}
