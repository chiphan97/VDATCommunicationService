import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SearchUsersComponent } from './search-users/search-users.component';
import {NzAvatarModule, NzFormModule, NzGridModule, NzIconModule, NzInputModule, NzListModule, NzTypographyModule} from 'ng-zorro-antd';
import {FormsModule} from '@angular/forms';



@NgModule({
  declarations: [SearchUsersComponent],
  exports: [
    SearchUsersComponent
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
    NzAvatarModule
  ]
})
export class UserModule { }
