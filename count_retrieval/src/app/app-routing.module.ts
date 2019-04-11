import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";

import { CountsComponent } from "./pages/counts/counts.component";

const routes: Routes = [
  { path: "", redirectTo: "/counts", pathMatch: "full" },
  { path: "counts", component: CountsComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
