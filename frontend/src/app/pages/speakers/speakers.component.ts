import { Component, OnInit, ViewChild } from "@angular/core";
import { SpeakerService } from "../../services/speaker.service";

import { Ng2SmartTableComponent } from "ng2-smart-table/ng2-smart-table.component";

import { TableSetting } from "../table_setting";
import {
  NumberInputComponent,
  TextRenderComponent,
  TextInputComponent
} from "../../shared_components";
import { LocalDataSource } from "ng2-smart-table";

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

  constructor(private speakerService: SpeakerService) {
    this.tableDataSource = new LocalDataSource();

    this.speakerService.getAllSpeakers().subscribe(data => {
      this.tableDataSource.load(data);
    });
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
    this.speakerService.writeSpeaker(speaker.firstName, speaker.lastName, speaker.email).subscribe(
      response => {
        speaker.id = response.id;
        event.confirm.resolve(speaker);
      },
      error => {
        console.log(error);
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
        event.confirm.reject();
      }
    );
  }

  deleteSpeaker(event): void {
    this.speakerService.deleteSpeaker(event.data.id).subscribe(
      () => { },
      error => {
        console.log(error);
      }
    );
    event.confirm.resolve();
  }
}
