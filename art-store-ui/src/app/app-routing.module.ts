import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AddArtComponent } from './add-art/add-art.component';
import { ListArtComponent } from './list-art/list-art.component';

const routes: Routes = [
  { path: 'add-art', component: AddArtComponent },
  { path: 'list-art', component: ListArtComponent },
  { path: '', redirectTo: 'add-art', pathMatch: 'full' } // default route
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
