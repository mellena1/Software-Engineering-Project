import { Component, OnInit, ViewChild } from "@angular/core";
import { SpeakerService } from "../../services/speaker.service";

import { Ng2SmartTableComponent } from "ng2-smart-table/ng2-smart-table.component";

import { TableSetting } from "../table_setting";
import {
  TextRenderComponent,
  TextInputComponent
} from "../../shared_components";
import { LocalDataSource } from "ng2-smart-table";
import { ErrorGlobals } from "src/app/globals/errors.global";

@Component({
  selector: "app-speakers",
  templateUrl: "./speakers.component.html"
})
export class SpeakersComponent implements OnInit {
  @ViewChild("table") table: Ng2SmartTableComponent;
  tableDataSource: LocalDataSource;

  columns = {
    firstName: {
      title: "First Name",
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "custom",
        component: TextInputComponent
      }
    },
    lastName: {
      title: "Last Name",
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "custom",
        component: TextInputComponent
      }
    },
    email: {
      title: "Email",
      type: "custom",
      renderComponent: TextRenderComponent,
      editor: {
        type: "custom",
        component: TextInputComponent
      }
    }
  };
  tableSettings = new TableSetting(this.columns);

  constructor(private speakerService: SpeakerService, private errorGlobals: ErrorGlobals) {
    this.tableDataSource = new LocalDataSource();

    this.speakerService.getAllSpeakers().subscribe(
      data => {
        this.tableDataSource.load(data);
      },
      error => {
        console.log(error);
      }
    );
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

  addASpeaker(event): void {
    var speaker = event.newData;
    this.speakerService
      .writeSpeaker(speaker.firstName, speaker.lastName, speaker.email)
      .subscribe(
        response => {
          speaker.id = response.id;
          event.confirm.resolve(speaker);
        },
        error => {
          console.log(error);
          if (error.status === 503) {
            this.errorGlobals.newError("The server is unavailable, please wait a minute and try again")
          } else {
            this.errorGlobals.newError("You must set one of the fields to add a speaker. The email address and names must also be valid");
          }
          event.confirm.reject();
        }
      );
  }

  updateSpeaker(event): void {
    var speaker = event.newData;
    this.speakerService.updateSpeaker(speaker).subscribe(
      () => {
        event.confirm.resolve(speaker);
      },
      error => {
        console.log(error);
        if (error.status === 503) {
          this.errorGlobals.newError("The server is unavailable, please wait a minute and try again")
        } else {
          this.errorGlobals.newError("You must change one of the values and the email address and names must be valid");
        }
        event.confirm.reject();
      }
    );
  }

  deleteSpeaker(event): void {
    this.speakerService.deleteSpeaker(event.data.id).subscribe(
      () => {},
      error => {
        console.log(error);
        if (error.status === 503) {
          this.errorGlobals.newError("The server is unavailable, please wait a minute and try again")
        } else {
          this.errorGlobals.newError("Delete failed");
        }
        event.confirm.reject();
      }
    );
    event.confirm.resolve();
  }
}
