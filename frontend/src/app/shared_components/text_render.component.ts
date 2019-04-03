import { Component, Input, OnInit } from "@angular/core";

import { ViewCell } from "ng2-smart-table";

@Component({
  template: `
    <div style="font-size: 20px;">{{ this.value }}</div>
  `
})
export class TextRenderComponent implements ViewCell, OnInit {
  @Input() value: string;
  @Input() rowData: any;

  ngOnInit() {}
}
