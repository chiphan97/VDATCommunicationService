import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppRouting} from './app.routing';
import {AppComponent} from './app.component';
import {NzGridModule, NzIconModule} from 'ng-zorro-antd';
import {BrowserAnimationsModule, NoopAnimationsModule} from '@angular/platform-browser/animations';
import {HttpClientModule} from '@angular/common/http';
import {SocketService} from './service/socket.service';

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    NzIconModule,
    AppRouting,
    NzGridModule
  ],
  providers: [SocketService],
  bootstrap: [AppComponent]
})
export class AppModule {
}
