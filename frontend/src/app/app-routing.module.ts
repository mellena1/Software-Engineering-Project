import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { SessionsComponent }   from './pages/sessions/sessions.component';
import { RoomsComponent } from './pages/rooms/rooms.component';
import { SpeakersComponent } from './pages/speakers/speakers.component';
import { TimeslotsComponent } from './pages/timeslots/timeslots.component';

const routes: Routes = [
  { path: '', redirectTo: '/sessions', pathMatch: 'full' },
  { path: 'sessions', component: SessionsComponent },
  { path: 'rooms', component: RoomsComponent},
  { path: 'speakers', component: SpeakersComponent},
  { path: 'timeslots', component: TimeslotsComponent},
  

];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}
