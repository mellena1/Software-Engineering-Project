import { Component, Input, OnInit } from "@angular/core";

import { ViewCell } from "ng2-smart-table";

@Component({
  template: `
    <div class="table-text">{{ this.value }}</div>
  `,
  styleUrls: ["./table.components.css"]
})
export class TextRenderComponent implements ViewCell, OnInit {
  @Input() value: string;
  @Input() rowData: any;

  ngOnInit() {
    if (this.value == null) {
      this.value = "";
    }
  }
}
