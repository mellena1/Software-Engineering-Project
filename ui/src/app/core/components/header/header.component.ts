import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-header',
  template: `
    <nav class="navbar is-dark">
      <!-- menu -->
      <div class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item" routerLink="">Home</a>
          <a class="navbar-item" routerLink="contact">Contact</a>
          <a class="navbar-item" routerLink="users">Users</a>
        </div>
      </div>
    </nav>
  `,
  styles: []
})
export class HeaderComponent implements OnInit {
  constructor() {}
  ngOnInit() {}
}
