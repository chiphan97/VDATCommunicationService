import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ArticleComponent} from './article.component';
import {ArticleRouting} from './article.routing';


@NgModule({
  declarations: [ArticleComponent],
  imports: [
    CommonModule,
    ArticleRouting
  ]
})
export class ArticleModule {
}
