import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ArticleEditorComponent} from './article-editor.component';
import {ArticleEditorRouting} from './article-editor.routing';


@NgModule({
  declarations: [ArticleEditorComponent],
  imports: [
    CommonModule,
    ArticleEditorRouting
  ]
})
export class ArticleEditorModule {
}
