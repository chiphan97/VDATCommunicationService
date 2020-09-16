import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateNewGroupComponent } from './create-new-group/create-new-group.component';
import { AddMemberGroupComponent } from './add-member-group/add-member-group.component';
import {
  NzAvatarModule,
  NzButtonModule,
  NzCheckboxModule,
  NzFormModule,
  NzGridModule,
  NzInputModule,
  NzModalModule,
  NzSelectModule
} from 'ng-zorro-antd';
import {ReactiveFormsModule} from '@angular/forms';



@NgModule({
  declarations: [CreateNewGroupComponent, AddMemberGroupComponent],
  imports: [
    CommonModule,
    NzGridModule,
    NzFormModule,
    ReactiveFormsModule,
    NzInputModule,
    NzSelectModule,
    NzAvatarModule,
    NzCheckboxModule,
    NzButtonModule,
    NzModalModule
  ]
})
export class GroupModule { }
