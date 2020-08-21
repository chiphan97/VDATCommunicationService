import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MessengerComponent} from './messenger.component';
import {MessengerRouting} from './messenger.routing';
import {ChatModule} from '../../page/chat/chat.module';
import {
  NzAffixModule,
  NzBreadCrumbModule,
  NzButtonModule,
  NzFormModule,
  NzIconModule,
  NzInputModule,
  NzLayoutModule,
  NzMenuModule
} from 'ng-zorro-antd';
import {SharedModule} from '../../shared/shared.module';
import {UserModule} from '../../component/user/user.module';


@NgModule({
  declarations: [MessengerComponent],
  imports: [
    CommonModule,
    ChatModule,
    MessengerRouting,
    NzLayoutModule,
    NzMenuModule,
    NzIconModule,
    SharedModule,
    UserModule,
    NzAffixModule,
    NzFormModule,
    NzInputModule,
    NzButtonModule
  ]
})
export class MessengerModule {
}
