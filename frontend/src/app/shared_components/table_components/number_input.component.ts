import { Component, OnInit } from "@angular/core";

import { DefaultEditor } from "ng2-smart-table";

@Component({
  template: `
    <input
      type="number"
      [(ngModel)]="stringNumber"
      class="table-text"
    />
    <div [hidden]="true" [innerHTML]="parseToInt()" #htmlValue></div>
  `,
  styleUrls: ["./table.components.css"]
})
export class NumberInputComponent extends DefaultEditor implements OnInit {
  stringNumber: string;

  ngOnInit() {
    this.stringNumber = this.cell.getValue();
  }

  parseToInt() {
    this.cell.newValue = parseInt(this.stringNumber);
  }
}
