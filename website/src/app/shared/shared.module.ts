import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import { FooterComponent } from './footer/footer.component';
import {NzLayoutModule} from 'ng-zorro-antd';



@NgModule({
  declarations: [NavBarComponent, FooterComponent],
  exports: [
    FooterComponent
  ],
  imports: [
    CommonModule,
    NzLayoutModule
  ]
})
export class SharedModule { }
