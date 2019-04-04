import { Injectable } from "@angular/core";

import * as moment from "moment";

@Injectable()
export class TimeslotGlobals {
  twelveHourIsChecked: boolean = true;

  formatTime(time: string): string {
    if (this.twelveHourIsChecked) {
      return this.format12Hour(time);
    } else {
      return this.format24Hour(time);
    }
  }

  format12Hour(time: string): string {
    return moment(time, "YYYY-MM-DDTHH:mm:ss").format("h:mm a");
  }

  format24Hour(time: string): string {
    return moment(time, "YYYY-MM-DDTHH:mm:ss").format("H:mm");
  }
}
