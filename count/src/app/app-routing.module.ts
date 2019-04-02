import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";

import { CountsComponent } from "./pages/counts/counts.component";


const routes: Routes = [
  { path: "", redirectTo: "/count", pathMatch: "full" },
  { path: "count", component: CountsComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
