import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MessengerContentComponent} from './messenger-content/messenger-content.component';
import {MessengerOptionComponent} from './messenger-option/messenger-option.component';
import {ScrollingModule} from '@angular/cdk/scrolling';
import {
  NzAffixModule, NzAvatarModule, NzBadgeModule, NzButtonModule, NzCollapseModule,
  NzCommentModule, NzDrawerModule, NzDropDownModule,
  NzFormModule,
  NzGridModule,
  NzIconModule,
  NzInputModule,
  NzListModule, NzModalModule, NzPageHeaderModule, NzPopconfirmModule,
  NzSkeletonModule, NzSwitchModule, NzToolTipModule
} from 'ng-zorro-antd';
import { MessengerHeaderComponent } from './messenger-header/messenger-header.component';
import { MessengerDrawerComponent } from './messenger-drawer/messenger-drawer.component';
import {NzResizableModule} from 'ng-zorro-antd/resizable';
import {FormsModule} from '@angular/forms';
import {NzSpaceModule} from 'ng-zorro-antd/space';
import { MessageSidebarLeftComponent } from './message-sidebar-left/message-sidebar-left.component';
import { MessageSidebarRightComponent } from './message-sidebar-right/message-sidebar-right.component';


@NgModule({
  declarations: [MessengerContentComponent, MessengerOptionComponent, MessengerHeaderComponent, MessengerDrawerComponent, MessageSidebarLeftComponent, MessageSidebarRightComponent],
  exports: [
    MessengerContentComponent,
    MessageSidebarLeftComponent,
    MessageSidebarRightComponent,
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
    NzDrawerModule,
    NzResizableModule,
    NzPopconfirmModule,
    NzModalModule,
    NzToolTipModule,
    NzDropDownModule,
    NzSwitchModule,
    FormsModule,
    NzSpaceModule,
    NzCollapseModule
  ]
})
export class ComponentModule {
}
