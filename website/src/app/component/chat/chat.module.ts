import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ChatSidebarLeftComponent} from './chat-sidebar-left/chat-sidebar-left.component';
import {ChatSidebarRightComponent} from './chat-sidebar-right/chat-sidebar-right.component';
import {ChatContentComponent} from './chat-content/chat-content.component';
import {ChatHeaderComponent} from './chat-header/chat-header.component';
import {ScrollingModule} from '@angular/cdk/scrolling';
import {
  NzAffixModule, NzAvatarModule, NzBadgeModule, NzButtonModule, NzCheckboxModule, NzCollapseModule,
  NzCommentModule, NzDrawerModule, NzDropDownModule,
  NzFormModule,
  NzGridModule,
  NzIconModule,
  NzInputModule,
  NzListModule, NzMentionModule, NzMessageModule, NzModalModule, NzPageHeaderModule, NzPopconfirmModule, NzSelectModule,
  NzSkeletonModule, NzSwitchModule, NzToolTipModule, NzTypographyModule
} from 'ng-zorro-antd';
import {NzResizableModule} from 'ng-zorro-antd/resizable';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {NzSpaceModule} from 'ng-zorro-antd/space';
import {GroupModule} from '../group/group.module';
import {MessageComponent} from './chat-content/message/message.component';

@NgModule({
  declarations: [
    ChatSidebarLeftComponent,
    ChatSidebarRightComponent,
    ChatContentComponent,
    ChatHeaderComponent,
    MessageComponent
  ],
  exports: [
    ChatSidebarLeftComponent,
    ChatHeaderComponent,
    ChatContentComponent,
    ChatSidebarRightComponent
  ],
  imports: [
    CommonModule,
    ScrollingModule,
    NzListModule,
    NzSkeletonModule,
    NzGridModule,
    NzFormModule,
    NzInputModule,
    NzIconModule,
    NzAffixModule,
    NzCommentModule,
    NzAvatarModule,
    NzPageHeaderModule,
    NzButtonModule,
    NzBadgeModule,
    NzResizableModule,
    NzPopconfirmModule,
    NzModalModule,
    NzToolTipModule,
    NzDropDownModule,
    NzSwitchModule,
    FormsModule,
    NzSpaceModule,
    NzCollapseModule,
    NzTypographyModule,
    NzModalModule,
    NzMentionModule,
    NzCheckboxModule,
    ReactiveFormsModule,
    NzSelectModule,
    NzMessageModule,
    GroupModule
  ]
})
export class ChatModule {
}
