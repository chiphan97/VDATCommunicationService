import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {AuthComponent} from './auth.component';
import {AuthRouting} from './auth.routing';
import {NzGridModule} from 'ng-zorro-antd';


@NgModule({
  declarations: [AuthComponent],
  imports: [
    CommonModule,
    AuthRouting,
    NzGridModule
  ]
})
export class AuthModule {
}
