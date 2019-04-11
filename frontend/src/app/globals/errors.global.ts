import { Injectable } from "@angular/core";
import { Subject } from "rxjs";

@Injectable()
export class ErrorGlobals {
  error = new Subject<string>();

  newError(err: string) {
    this.error.next(err);
  }
}
