import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SearchUsersComponent } from './search-users/search-users.component';
import {NzFormModule, NzGridModule, NzIconModule, NzInputModule, NzListModule, NzTypographyModule} from 'ng-zorro-antd';



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
    NzTypographyModule
  ]
})
export class UserModule { }
