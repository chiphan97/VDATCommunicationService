import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MessengerComponent} from './messenger.component';
import {MessengerRouting} from './messenger.routing';
import {MessengerSidebarLeftComponent} from './messenger-sidebar-left/messenger-sidebar-left.component';
import {MessengerSidebarRightComponent} from './messenger-sidebar-right/messenger-sidebar-right.component';
import {MessengerHeaderComponent} from './messenger-header/messenger-header.component';
import {MessengerContentComponent} from './messenger-content/messenger-content.component';
import {
  NzAvatarModule,
  NzBadgeModule,
  NzButtonModule,
  NzCollapseModule,
  NzCommentModule,
  NzDropDownModule,
  NzFormModule,
  NzGridModule,
  NzIconModule,
  NzInputModule,
  NzListModule,
  NzMessageModule,
  NzModalModule,
  NzPageHeaderModule,
  NzPopconfirmModule,
  NzPopoverModule,
  NzResultModule,
  NzSpinModule,
  NzSwitchModule,
  NzToolTipModule,
  NzTypographyModule,
  NzUploadModule
} from 'ng-zorro-antd';
import {NzResizableModule} from 'ng-zorro-antd/resizable';
import {NzSpaceModule} from 'ng-zorro-antd/space';
import {MessengerMessageComponent} from './messenger-message/messenger-message.component';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {SettingModule} from '../../component/setting/setting.module';
import {GroupModule} from '../../component/group/group.module';
import { MessengerReplyThreadRightComponent } from './messenger-reply-thread-right/messenger-reply-thread-right.component';
import { ReplyThreadHeaderComponent } from './messenger-reply-thread-right/reply-thread-header/reply-thread-header.component';

@NgModule({
  declarations: [
    MessengerComponent,
    MessengerSidebarLeftComponent,
    MessengerSidebarRightComponent,
    MessengerHeaderComponent,
    MessengerContentComponent,
    MessengerMessageComponent,
    MessengerReplyThreadRightComponent,
    ReplyThreadHeaderComponent,
  ],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NzGridModule,
    NzResizableModule,
    NzButtonModule,
    NzDropDownModule,
    NzToolTipModule,
    NzIconModule,
    NzSpaceModule,
    NzSwitchModule,
    NzInputModule,
    NzFormModule,
    NzListModule,
    NzBadgeModule,
    NzAvatarModule,
    NzTypographyModule,
    NzPopoverModule,
    NzCollapseModule,
    NzPopconfirmModule,
    NzModalModule,
    NzMessageModule,
    NzPageHeaderModule,
    NzResultModule,
    NzCommentModule,
    SettingModule,
    GroupModule,
    MessengerRouting,
    NzSpinModule,
    NzUploadModule
  ]
})
export class MessengerModule {
}
