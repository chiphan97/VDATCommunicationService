import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {AuthComponent} from './auth.component';
import {AuthRouting} from './auth.routing';
import {NzAvatarModule, NzButtonModule, NzCardModule, NzGridModule, NzIconModule, NzLayoutModule, NzTypographyModule} from 'ng-zorro-antd';


@NgModule({
  declarations: [AuthComponent],
  imports: [
    CommonModule,
    AuthRouting,
    NzLayoutModule,
    NzGridModule,
    NzTypographyModule,
    NzCardModule,
    NzAvatarModule,
    NzButtonModule,
    NzIconModule
  ]
})
export class AuthModule {
}
