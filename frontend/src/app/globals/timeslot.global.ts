import { Injectable } from "@angular/core";

import * as moment from "moment";

@Injectable()
export class TimeslotGlobals {
  twelveHourIsChecked: boolean = true;

  formatTime(time: string): string {
    if (this.twelveHourIsChecked) {
      return TimeslotGlobals.format12Hour(time);
    } else {
      return TimeslotGlobals.format24Hour(time);
    }
  }

  static isValidTime(time: string): boolean {
    return moment(time, "YYYY-MM-DDTHH:mm:ss").isValid();
  }

  static sortTime(timeA: string, timeB: string): number {
    return (
      moment(timeA, "YYYY-MM-DDTHH:mm:ss").valueOf() -
      moment(timeB, "YYYY-MM-DDTHH:mm:ss").valueOf()
    );
  }

  static format12Hour(time: string): string {
    return moment(time, "YYYY-MM-DDTHH:mm:ss").format("h:mm a");
  }

  static format24Hour(time: string): string {
    return moment(time, "YYYY-MM-DDTHH:mm:ss").format("H:mm");
  }
}
