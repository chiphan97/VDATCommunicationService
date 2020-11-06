import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SearchUsersComponent } from './search-users/search-users.component';
import {
  NzAvatarModule,
  NzButtonModule, NzCheckboxModule,
  NzFormModule,
  NzGridModule,
  NzIconModule,
  NzInputModule,
  NzListModule, NzModalModule, NzSkeletonModule,
  NzTypographyModule
} from 'ng-zorro-antd';
import {FormsModule} from '@angular/forms';
import {ScrollingModule} from '@angular/cdk/scrolling';


@NgModule({
  declarations: [SearchUsersComponent],
  exports: [
    SearchUsersComponent,
  ],
  imports: [
    CommonModule,
    NzGridModule,
    NzFormModule,
    NzInputModule,
    NzIconModule,
    NzListModule,
    NzTypographyModule,
    FormsModule,
    NzAvatarModule,
    NzButtonModule,
    NzModalModule,
    ScrollingModule,
    NzSkeletonModule,
    NzCheckboxModule
  ]
})
export class UserModule { }
