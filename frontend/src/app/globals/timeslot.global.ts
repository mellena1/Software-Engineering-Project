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
    // Ignore date
    var a = moment(timeA, "YYYY-MM-DDTHH:mm:ss").month(0).date(0).year(0);
    var b = moment(timeB, "YYYY-MM-DDTHH:mm:ss").month(0).date(0).year(0);
    return a.isBefore(b) ? -1 : 1;
  }

  static format12Hour(time: string): string {
    return moment(time, "YYYY-MM-DDTHH:mm:ss").format("h:mm a");
  }

  static format24Hour(time: string): string {
    return moment(time, "YYYY-MM-DDTHH:mm:ss").format("H:mm");
  }
}
