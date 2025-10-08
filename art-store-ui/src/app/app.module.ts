import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AddArtComponent } from './add-art/add-art.component';
import { ListArtComponent } from './list-art/list-art.component';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';


@NgModule({
  declarations: [
    AppComponent,
    AddArtComponent,
    ListArtComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
   HttpClientModule,
  FormsModule  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
