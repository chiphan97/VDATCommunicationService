import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {AuthComponent} from './auth.component';
import {MessengerModule} from '../../page/messenger/messenger.module';
import {AuthRouting} from './auth.routing';
import {NzGridModule, NzLayoutModule} from 'ng-zorro-antd';
import {SharedModule} from '../../shared/shared.module';
import {NzResizableModule} from 'ng-zorro-antd/resizable';
import {ComponentModule} from '../../component/component.module';


@NgModule({
  declarations: [AuthComponent],
  imports: [
    CommonModule,
    MessengerModule,
    AuthRouting,
    NzLayoutModule,
    SharedModule,
    NzGridModule,
    NzResizableModule,
    ComponentModule
  ]
})
export class AuthModule {
}
