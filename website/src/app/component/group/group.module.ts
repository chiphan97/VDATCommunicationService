import {NgModule, NO_ERRORS_SCHEMA} from '@angular/core';
import {CommonModule} from '@angular/common';
import {CreateNewGroupComponent} from './create-new-group/create-new-group.component';
import {AddMemberGroupComponent} from './add-member-group/add-member-group.component';
import {
  NzAvatarModule,
  NzButtonModule,
  NzCheckboxModule,
  NzFormModule,
  NzGridModule, NzIconModule,
  NzInputModule,
  NzModalModule,
  NzSelectModule, NzTableModule
} from 'ng-zorro-antd';
import {ReactiveFormsModule} from '@angular/forms';
import {UserModule} from '../user/user.module';


@NgModule({
  declarations: [
    CreateNewGroupComponent,
    AddMemberGroupComponent
  ],
  imports: [
    CommonModule,
    NzGridModule,
    UserModule,
    NzButtonModule,
    NzFormModule,
    NzInputModule,
    ReactiveFormsModule,
    NzCheckboxModule
  ],
  exports: [
    CreateNewGroupComponent,
    AddMemberGroupComponent
  ],
  schemas: [
    NO_ERRORS_SCHEMA
  ]
})
export class GroupModule {
}
