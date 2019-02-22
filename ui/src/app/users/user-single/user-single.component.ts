import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { UserService } from '../../core/services/user.service';

@Component({
  selector: 'app-user-single',
  template: `
    <section class="section">
    <div class="container">
  
      <div class="card" *ngIf="user">
        <img [src]="user.avatar_url">
        <h2>{{ user.login }}</h2>
      </div>

    </div>
    </section>
  `,
  styles: []
})
export class UserSingleComponent implements OnInit {
  user; // property to hold our user when we eventually grab them from github api

  // inject the things we need
  constructor(
    private route: ActivatedRoute,
    private userService: UserService
  ) {}

  // watch the route parameters for changes
  // every time it changes (or on first load), go get a user from userservice
  ngOnInit() {
    this.route.params.subscribe(params => {
      const username = params['username'];
      this.userService
          .getUser(username)
          .subscribe(user => this.user = user);
    });
  }
}
