import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { HttpClientModule } from "@angular/common/http";

import { HttpClientInMemoryWebApiModule } from "angular-in-memory-web-api";

import { AppRoutingModule } from "./app-routing.module";

import { AppComponent } from "./app.component";
import { CounterComponent } from "./pages/counter/counter.component";
import { LoginComponent } from "./pages/login/login.component";
// import {  } from "./pages/phase2sessions/phase2sessions.component";

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    ReactiveFormsModule,
    HttpClientModule
  ],
  declarations: [AppComponent, CounterComponent, LoginComponent],
  bootstrap: [AppComponent]
})
export class AppModule {}
