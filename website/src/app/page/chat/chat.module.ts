import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ChatComponent} from './chat.component';
import {NzButtonModule, NzCardModule, NzFormModule, NzGridModule, NzInputModule, NzListModule} from 'ng-zorro-antd';
import {UserModule} from '../../component/user/user.module';
import {NzSpaceModule} from 'ng-zorro-antd/space';
import {FormsModule} from '@angular/forms';


@NgModule({
  declarations: [ChatComponent],
  exports: [
    ChatComponent
  ],
  imports: [
    CommonModule,
    NzGridModule,
    UserModule,
    NzSpaceModule,
    NzCardModule,
    NzInputModule,
    NzButtonModule,
    NzFormModule,
    FormsModule,
    NzListModule
  ]
})
export class ChatModule {
}
