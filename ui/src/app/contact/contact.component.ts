import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-contact',
  template: `
    <section class="hero is-primary is-bold">
    <div class="hero-body">
    <div class="container">  
      <h1 class="title">Contact Us!</h1>  
    </div>
    </div>
    </section>

    <section class="section">
    <div class="container">

      <!-- form goes here -->
      <form (ngSubmit)="processForm()">

        <!-- name -->
        <div class="field">
          <input 
            type="text" 
            name="name" 
            class="input" 
            placeholder="Your Name" 
            [(ngModel)]="name"
            required
            minlength="3"
            #nameInput="ngModel">

          <div class="help is-error" *ngIf="nameInput.invalid && nameInput.dirty">
            Name is required and needs to be at least 3 characters long.
          </div>
        </div>

        <!-- email -->
        <div class="field">          
          <input 
            type="email" 
            name="email" 
            class="input" 
            placeholder="Your Email" 
            [(ngModel)]="email"
            required
            email
            #emailInput="ngModel">

          <div class="help is-error" *ngIf="emailInput.invalid && emailInput.dirty">
            Needs to be a valid email.
          </div>
        </div>

        <!-- message -->
        <div class="field">
          <textarea 
            class="textarea" 
            name="message" 
            placeholder="What's on your mind?" 
            [(ngModel)]="message"
            required
            #messageInput="ngModel"></textarea>

            <div class="help is-error" *ngIf="emailInput.invalid && emailInput.dirty">
              Your message is required!
            </div>
        </div>

        <button type="submit" class="button is-danger is-large">Submit</button>

      </form>

    </div>
    </section>
  `,
  styles: []
})
export class ContactComponent implements OnInit {
  name: string;
  email: string;
  message: string;

  constructor() {}

  ngOnInit() {}

  /**
   * Process the form we have. Send to whatever backend
   * Only alerting for now
   */
  processForm() {
    const allInfo = `My name is ${this.name}. My email is ${this.email}. My message is ${this.message}`;
    alert(allInfo); 
  }

}
