import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MasterComponent} from './master.component';
import {MasterRouting} from './master.routing';
import {NzLayoutModule} from 'ng-zorro-antd';
import {SharedModule} from '../../shared/shared.module';


@NgModule({
  declarations: [MasterComponent],
  imports: [
    CommonModule,
    NzLayoutModule,
    SharedModule,
    MasterRouting,
  ]
})
export class MasterModule {
}
