import { Component, OnInit, ViewChild } from "@angular/core";
import { TimeslotService } from "src/app/services/timeslot.service";
import { TimeslotGlobals } from "../../globals/timeslot.global";

import { Ng2SmartTableComponent } from "ng2-smart-table/ng2-smart-table.component";

import { TableSetting } from "../table_setting";
import {
  TimeslotInputComponent,
  TimeslotRenderComponent
} from "../../shared_components";
import { LocalDataSource } from "ng2-smart-table";
import { ErrorGlobals } from "src/app/globals/errors.global";

@Component({
  selector: "app-timeslots",
  templateUrl: "./timeslots.component.html"
})
export class TimeslotsComponent implements OnInit {
  @ViewChild("table") table: Ng2SmartTableComponent;
  tableDataSource: LocalDataSource;

  columns = {
    startTime: {
      title: "Start Time",
      type: "custom",
      renderComponent: TimeslotRenderComponent,
      editor: {
        type: "custom",
        component: TimeslotInputComponent
      }
    },
    endTime: {
      title: "End Time",
      type: "custom",
      renderComponent: TimeslotRenderComponent,
      editor: {
        type: "custom",
        component: TimeslotInputComponent
      }
    }
  };
  tableSettings = new TableSetting(this.columns);

  constructor(
    private timeslotService: TimeslotService,
    private timeslotGlobals: TimeslotGlobals,
    private errorGlobals: ErrorGlobals
  ) {
    this.tableDataSource = new LocalDataSource();

    this.getAllTimeslots();
  }

  ngOnInit() {
    this.table.userRowSelect.subscribe(() => {
      this.table.grid.dataSet.deselectAll();
    });

    this.table.initGrid();

    this.tableDataSource.onChanged().subscribe(() => {
      this.table.grid.createFormShown = true;
    });
  }

  getAllTimeslots() {
    this.timeslotService.getAllTimeslots().subscribe(
      data => {
        data.sort((a, b) => {
          return TimeslotGlobals.sortTime(a.startTime, b.startTime);
        });
        this.tableDataSource.load(data);
      },
      error => {
        console.log(error);
      }
    );
  }

  addATimeslot(event): void {
    var timeslot = event.newData;

    this.timeslotService
      .writeTimeslot(timeslot.startTime, timeslot.endTime)
      .subscribe(
        response => {
          timeslot.id = response.id;
          event.confirm.resolve(timeslot);
        },
        error => {
          console.log(error);
          if (error.status === 503) {
            this.errorGlobals.newError(
              "The server is unavailable, please wait a minute and try again"
            );
          } else {
            this.errorGlobals.newError(
              "You must set both times to add a timeslot"
            );
          }
          event.confirm.reject();
        }
      );
  }

  updateTimeslot(event): void {
    var timeslot = event.newData;
    this.timeslotService.updateTimeslot(timeslot).subscribe(
      () => {
        event.confirm.resolve(timeslot);
      },
      error => {
        console.log(error);
        if (error.status === 503) {
          this.errorGlobals.newError(
            "The server is unavailable, please wait a minute and try again"
          );
        } else {
          this.errorGlobals.newError(
            "You must change a field and set both times"
          );
        }
        event.confirm.reject();
      }
    );
  }

  deleteTimeslot(event): void {
    this.timeslotService.deleteTimeslot(event.data.id).subscribe(
      () => {
        event.confirm.resolve();
      },
      error => {
        console.log(error);
        if (error.status === 503) {
          this.errorGlobals.newError(
            "The server is unavailable, please wait a minute and try again"
          );
        } else {
          this.errorGlobals.newError("Delete failed");
        }
        event.confirm.reject();
      }
    );
  }
}
