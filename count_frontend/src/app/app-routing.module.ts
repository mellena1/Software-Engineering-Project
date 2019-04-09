import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";

import { CountsComponent } from "./pages/counts/counts.component";
import { LoginComponent } from "./pages/login/login.component";

const routes: Routes = [
  { path: "", redirectTo: "/login", pathMatch: "full" },
  { path: "login", component: LoginComponent },
  { path: "count", component: CountsComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
