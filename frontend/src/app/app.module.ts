import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { HttpClientModule } from "@angular/common/http";

import { Ng2SmartTableModule } from "ng2-smart-table";
import { NgbTimepickerModule, NgbAlertModule } from "@ng-bootstrap/ng-bootstrap";

import { AppRoutingModule } from "./app-routing.module";

import { AppComponent } from "./app.component";
import { SessionsComponent } from "./pages/sessions/sessions.component";
import { RoomsComponent } from "./pages/rooms/rooms.component";
import { SpeakersComponent } from "./pages/speakers/speakers.component";
import { TimeslotsComponent } from "./pages/timeslots/timeslots.component";

import {
  TwelveTwentyfourHourRadioComponent,
  ErrorComponent,
  NumberInputComponent,
  TextRenderComponent,
  TextInputComponent,
  TextListInputComponent,
  TimeslotRenderComponent,
  TimeslotInputComponent,
  TimeslotListInputComponent
} from "./shared_components";
import { TimeslotGlobals } from "./globals/timeslot.global";
import { ErrorGlobals } from "./globals/errors.global";

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    ReactiveFormsModule,
    HttpClientModule,
    Ng2SmartTableModule,
    NgbTimepickerModule,
    NgbAlertModule
  ],
  declarations: [
    AppComponent,
    SessionsComponent,
    RoomsComponent,
    SpeakersComponent,
    TimeslotsComponent,
    NumberInputComponent,
    TextInputComponent,
    TextRenderComponent,
    TextListInputComponent,
    TimeslotRenderComponent,
    TimeslotInputComponent,
    TimeslotListInputComponent,
    TwelveTwentyfourHourRadioComponent,
    ErrorComponent
  ],
  entryComponents: [
    NumberInputComponent,
    TextInputComponent,
    TextRenderComponent,
    TextListInputComponent,
    TimeslotRenderComponent,
    TimeslotInputComponent,
    TimeslotListInputComponent,
    TwelveTwentyfourHourRadioComponent,
    ErrorComponent
  ],
  providers: [TimeslotGlobals, ErrorGlobals],
  bootstrap: [AppComponent]
})
export class AppModule {}
