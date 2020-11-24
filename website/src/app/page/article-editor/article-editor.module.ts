import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ArticleEditorComponent} from './article-editor.component';
import {ArticleEditorRouting} from './article-editor.routing';
import {EditorModule} from '../../component/editor/editor.module';
import {NzFormModule, NzGridModule, NzIconModule, NzInputModule, NzSelectModule} from 'ng-zorro-antd';
import {ReactiveFormsModule} from '@angular/forms';


@NgModule({
  declarations: [ArticleEditorComponent],
  imports: [
    CommonModule,
    ArticleEditorRouting,
    EditorModule,
    NzGridModule,
    NzIconModule,
    NzInputModule,
    NzFormModule,
    ReactiveFormsModule,
    NzSelectModule
  ]
})
export class ArticleEditorModule {
}
