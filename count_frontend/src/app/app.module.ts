import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { HttpClientModule } from "@angular/common/http";

import { AppRoutingModule } from "./app-routing.module";

import { AppComponent } from "./app.component";
import { CountComponent } from "./pages/count/count.component";
import { LoginComponent } from "./pages/login/login.component";

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    ReactiveFormsModule,
    HttpClientModule
  ],
  declarations: [AppComponent, CountComponent, LoginComponent],
  bootstrap: [AppComponent]
})
export class AppModule {}
