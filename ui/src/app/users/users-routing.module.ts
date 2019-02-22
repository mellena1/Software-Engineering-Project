import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { UsersListComponent } from './users-list/users-list.component';
import { UserSingleComponent } from './user-single/user-single.component';

const routes: Routes = [
  {
    path: '',
    component: UsersListComponent
  },
  {
    path: ':username',
    component: UserSingleComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UsersRoutingModule { }
