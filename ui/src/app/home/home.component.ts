import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-home',
  template: `
    <section class="hero is-info is-fullheight is-bold">
    <div class="hero-body">
    <div class="container">

      <h1 class="title">Home Page</h1>

    </div>
    </div>
    </section>
  `,
  styles: [`
    .hero {
      background-size: cover;
      background-position: center center;
    }
  `]
})
export class HomeComponent implements OnInit {
  constructor() {}

  ngOnInit() {}
}
