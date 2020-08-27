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
import {MessengerSidebarComponent} from './messenger-sidebar/messenger-sidebar.component';
import { MessengerHeaderComponent } from './messenger-header/messenger-header.component';
import { MessengerDrawerComponent } from './messenger-drawer/messenger-drawer.component';
import {NzResizableModule} from 'ng-zorro-antd/resizable';
import {FormsModule} from '@angular/forms';
import {NzSpaceModule} from 'ng-zorro-antd/space';


@NgModule({
  declarations: [MessengerContentComponent, MessengerOptionComponent, MessengerSidebarComponent, MessengerHeaderComponent, MessengerDrawerComponent],
  exports: [
    MessengerContentComponent,
    MessengerSidebarComponent
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
