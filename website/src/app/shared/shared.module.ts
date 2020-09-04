import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import { FooterComponent } from './footer/footer.component';
import {
  NzAvatarModule,
  NzButtonModule,
  NzDropDownModule,
  NzGridModule,
  NzIconModule,
  NzLayoutModule,
  NzTypographyModule
} from 'ng-zorro-antd';
import {NzSpaceModule} from 'ng-zorro-antd/space';



@NgModule({
  declarations: [NavBarComponent, FooterComponent],
    exports: [
        FooterComponent,
        NavBarComponent
    ],
  imports: [
    CommonModule,
    NzLayoutModule,
    NzGridModule,
    NzButtonModule,
    NzIconModule,
    NzAvatarModule,
    NzDropDownModule,
    NzSpaceModule,
    NzTypographyModule
  ]
})
export class SharedModule { }
