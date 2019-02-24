import { NgModule }       from '@angular/core';
import { BrowserModule }  from '@angular/platform-browser';
import { FormsModule }    from '@angular/forms';
import { HttpClientModule }    from '@angular/common/http';

import { HttpClientInMemoryWebApiModule } from 'angular-in-memory-web-api';

import { AppRoutingModule }     from './app-routing.module';

import { AppComponent }         from './app.component';
import { DashboardComponent }   from './dashboard/dashboard.component';
import { RoomsComponent } from './rooms/rooms.component';
import { SpeakersComponent } from './speakers/speakers.component';
import { TimeslotsComponent } from './timeslots/timeslots.component';

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
    TimeslotsComponent
  ],
  bootstrap: [ AppComponent ]
})
export class AppModule { }
