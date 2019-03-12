import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { HttpClientModule } from "@angular/common/http";

import { HttpClientInMemoryWebApiModule } from "angular-in-memory-web-api";

import { AppRoutingModule } from "./app-routing.module";

import { AppComponent } from "./app.component";
import { SessionsComponent } from "./pages/sessions/sessions.component";
import { RoomsComponent } from "./pages/rooms/rooms.component";
import { SpeakersComponent } from "./pages/speakers/speakers.component";
import { TimeslotsComponent } from "./pages/timeslots/timeslots.component";

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    ReactiveFormsModule,
    HttpClientModule
  ],
  declarations: [
    AppComponent,
    SessionsComponent,
    RoomsComponent,
    SpeakersComponent,
    TimeslotsComponent
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}
