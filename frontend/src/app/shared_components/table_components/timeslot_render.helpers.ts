import { TimeslotGlobals } from "../../globals/timeslot.global";

export class TimeslotRenderHelpers {
    static formatTimeForCell(time: string, timeslotGlobals: TimeslotGlobals): string {
        if (time.includes(" ")) {
          return TimeslotRenderHelpers.formatTimeForCellStartDashEnd(time, timeslotGlobals);
        } else {
          return TimeslotRenderHelpers.formatTimeForCellSingleTime(time, timeslotGlobals);
        }
      }
    
      static formatTimeForCellStartDashEnd(time: string, timeslotGlobals: TimeslotGlobals): string {
        var splitValue = time.split(" ");
        var startTime = splitValue[0];
        var endTime = splitValue[1];
    
        if (TimeslotGlobals.isValidTime(startTime) && TimeslotGlobals.isValidTime(endTime)) {
          return `${TimeslotRenderHelpers.formatTimeForCellSingleTime(
            startTime, timeslotGlobals
          )}-${TimeslotRenderHelpers.formatTimeForCellSingleTime(endTime, timeslotGlobals)}`;
        } else {
          // will return whatever is valid, or '' if neither are valid
          return `${TimeslotRenderHelpers.formatTimeForCellSingleTime(
            startTime, timeslotGlobals
          )}${TimeslotRenderHelpers.formatTimeForCellSingleTime(endTime, timeslotGlobals)}`
        }
      }
    
      static formatTimeForCellSingleTime(time: string, timeslotGlobals: TimeslotGlobals): string {
        if (TimeslotGlobals.isValidTime(time)) {
          // timeslotGlobals holds whether 12 or 24 hour time is selected
          return `${timeslotGlobals.formatTime(time)}`;
        } else {
          return '';
        }
      }
}
