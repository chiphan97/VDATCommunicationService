import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ListUserChatComponent } from './list-user-chat/list-user-chat.component';
import {NzDividerModule, NzGridModule, NzIconModule, NzListModule, NzMenuModule} from 'ng-zorro-antd';
import {NzTransitionPatchModule} from 'ng-zorro-antd/core/transition-patch/transition-patch.module';



@NgModule({
  declarations: [ListUserChatComponent],
  exports: [
    ListUserChatComponent
  ],
  imports: [
    CommonModule,
    NzMenuModule,
    NzIconModule,
    NzListModule,
    NzGridModule,
    NzDividerModule
  ]
})
export class UserModule { }
