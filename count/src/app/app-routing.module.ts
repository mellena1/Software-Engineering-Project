import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";

import { CounterComponent } from "./pages/counter/counter.component";
import { LoginComponent } from "./pages/login/login.component";
// import {  } from "./pages/phase2sessions/phase2sessions.component";

const routes: Routes = [
  { path: "", redirectTo: "/login", pathMatch: "full" },
  { path: "counter", component: CounterComponent },
  { path: "login", component: LoginComponent }
  // { path: "sessions", component:  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
