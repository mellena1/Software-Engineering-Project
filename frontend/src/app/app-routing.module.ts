import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { DashboardComponent }   from './dashboard/dashboard.component';
import { RoomsComponent } from './rooms/rooms.component';
import { SpeakersComponent } from './speakers/speakers.component';
import { TimeslotsComponent } from './timeslots/timeslots.component';

const routes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  { path: 'dashboard', component: DashboardComponent },
  { path: 'rooms', component: RoomsComponent},
  { path: 'speakers', component: SpeakersComponent},
  { path: 'timeslots', component: TimeslotsComponent}

];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}
