import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { EditorComponent } from './editor.component';
import {NzCardModule, NzGridModule} from 'ng-zorro-antd';



@NgModule({
    declarations: [EditorComponent],
    exports: [
        EditorComponent
    ],
  imports: [
    CommonModule,
    NzCardModule,
    NzGridModule
  ]
})
export class EditorModule { }
