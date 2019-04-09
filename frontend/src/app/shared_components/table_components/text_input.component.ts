import { Component, OnInit } from "@angular/core";

import { DefaultEditor } from "ng2-smart-table";

@Component({
  template: `
    <input [(ngModel)]="this.cell.newValue" class="table-text" />
  `,
  styleUrls: ["./table.components.css"]
})
export class TextInputComponent extends DefaultEditor implements OnInit {
  ngOnInit() {
      this.cell.newValue = this.cell.getValue();
  }
}
