import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { HttpClientModule } from "@angular/common/http";

import { Ng2SmartTableModule } from 'ng2-smart-table';

import { AppRoutingModule } from "./app-routing.module";

import { AppComponent } from "./app.component";
import { SessionsComponent } from "./pages/sessions/sessions.component";
import { RoomsComponent } from "./pages/rooms/rooms.component";
import { SpeakersComponent } from "./pages/speakers/speakers.component";
import { TimeslotsComponent } from "./pages/timeslots/timeslots.component";

import { NumberInputComponent, TextRenderComponent, TextInputComponent } from "./shared_components"

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    ReactiveFormsModule,
    HttpClientModule,
    Ng2SmartTableModule
  ],
  declarations: [
    AppComponent,
    SessionsComponent,
    RoomsComponent,
    SpeakersComponent,
    TimeslotsComponent,
    NumberInputComponent,
    TextInputComponent,
    TextRenderComponent
  ],
  entryComponents: [
    NumberInputComponent,
    TextInputComponent,
    TextRenderComponent
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}
