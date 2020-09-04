import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MasterComponent} from './master.component';
import {MasterRouting} from './master.routing';
import {NzLayoutModule} from 'ng-zorro-antd';
import {SharedModule} from '../../shared/shared.module';
import {MessengerModule} from '../../page/messenger/messenger.module';


@NgModule({
  declarations: [MasterComponent],
  imports: [
    CommonModule,
    NzLayoutModule,
    SharedModule,
    MessengerModule,
    MasterRouting,
  ]
})
export class MasterModule {
}
