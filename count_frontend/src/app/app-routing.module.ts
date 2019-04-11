import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";

import { CountComponent } from "./pages/count/count.component";
import { LoginComponent } from "./pages/login/login.component";

const routes: Routes = [
  { path: "", redirectTo: "/login", pathMatch: "full" },
  { path: "login", component: LoginComponent },
  { path: "count", component: CountComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
