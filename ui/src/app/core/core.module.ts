import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

import { HeaderComponent } from './components/header/header.component';
import { UserService } from './services/user.service';

@NgModule({
  imports: [
    CommonModule,
    RouterModule
  ],
  providers: [
    UserService
  ],
  declarations: [
    HeaderComponent
  ],
  exports: [
    HeaderComponent
  ],
})
export class CoreModule { }
