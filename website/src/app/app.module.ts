import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRouting } from './app.routing';
import { AppComponent } from './app.component';
import {NzGridModule, NzIconModule} from 'ng-zorro-antd';
import {NoopAnimationsModule} from '@angular/platform-browser/animations';
import {HttpClientModule} from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    NoopAnimationsModule,
    HttpClientModule,
    NzIconModule,
    AppRouting,
    NzGridModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
