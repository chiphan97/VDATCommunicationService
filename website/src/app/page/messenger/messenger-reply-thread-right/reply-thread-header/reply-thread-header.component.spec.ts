import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ReplyThreadHeaderComponent } from './reply-thread-header.component';

describe('ReplyThreadHeaderComponent', () => {
  let component: ReplyThreadHeaderComponent;
  let fixture: ComponentFixture<ReplyThreadHeaderComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ReplyThreadHeaderComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ReplyThreadHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
