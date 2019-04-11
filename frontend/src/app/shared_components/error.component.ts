import { Component, OnInit } from "@angular/core";
import { ErrorGlobals } from "../globals/errors.global";
import { debounceTime } from "rxjs/operators";

@Component({
  selector: "error",
  template: `
    <ngb-alert
      *ngIf="errorMessage"
      type="danger"
      (close)="errorMessage = null"
      >{{ errorMessage }}</ngb-alert
    >
  `
})
export class ErrorComponent implements OnInit {
  staticAlertClosed = false;
  errorMessage: string;

  constructor(private errorGlobals: ErrorGlobals) {}

  ngOnInit() {
    this.errorGlobals.error.subscribe(message => (this.errorMessage = message));
    this.errorGlobals.error
      .pipe(debounceTime(10000))
      .subscribe(() => (this.errorMessage = null));
  }
}
