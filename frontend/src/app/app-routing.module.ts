import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { DashboardComponent }   from './pages/dashboard/dashboard.component';
import { RoomsComponent } from './pages/rooms/rooms.component';
import { SpeakersComponent } from './pages/speakers/speakers.component';
import { TimeslotsComponent } from './pages/timeslots/timeslots.component';

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
