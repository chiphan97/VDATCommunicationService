import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {AuthenticateComponent} from './authenticate/authenticate.component';
import {IntegratedComponent} from './integrated/integrated.component';
import {
  NzAvatarModule,
  NzButtonModule,
  NzCardModule,
  NzGridModule,
  NzIconModule,
  NzLayoutModule,
  NzSpinModule,
  NzTypographyModule
} from 'ng-zorro-antd';
import {AuthenticateRouting} from './authenticate.routing';


@NgModule({
  declarations: [
    AuthenticateComponent,
    IntegratedComponent
  ],
  imports: [
    CommonModule,
    NzLayoutModule,
    NzGridModule,
    NzTypographyModule,
    NzCardModule,
    NzAvatarModule,
    NzButtonModule,
    NzIconModule,
    AuthenticateRouting,
    NzSpinModule
  ]
})
export class AuthenticateModule {
}
