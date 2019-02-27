import { NgModule }       from '@angular/core';
import { BrowserModule }  from '@angular/platform-browser';
import { FormsModule }    from '@angular/forms';
import { HttpClientModule }    from '@angular/common/http';

import { HttpClientInMemoryWebApiModule } from 'angular-in-memory-web-api';

import { AppRoutingModule }     from './app-routing.module';

import { AppComponent }         from './app.component';
import { DashboardComponent }   from './pages/dashboard/dashboard.component';
import { RoomsComponent } from './pages/rooms/rooms.component';
import { SpeakersComponent } from './pages/speakers/speakers.component';
import { TimeslotsComponent } from './pages/timeslots/timeslots.component';
import { TimeformComponent } from './forms/timeform/timeform.component';
import { RoomformComponent } from './forms/roomform/roomform.component';
import { SpeakerformComponent } from './forms/speakerform/speakerform.component';
import { SessionformComponent } from './forms/sessionform/sessionform.component';

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    HttpClientModule
  ],
  declarations: [
    AppComponent,
    DashboardComponent,
    RoomsComponent,
    SpeakersComponent,
    TimeslotsComponent,
    TimeformComponent,
    RoomformComponent,
    SpeakerformComponent,
    SessionformComponent
  ],
  bootstrap: [ AppComponent ]
})
export class AppModule { }
