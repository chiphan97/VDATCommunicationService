import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MessengerReplyThreadRightComponent } from './messenger-reply-thread-right.component';

describe('MessengerReplyThreadRightComponent', () => {
  let component: MessengerReplyThreadRightComponent;
  let fixture: ComponentFixture<MessengerReplyThreadRightComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MessengerReplyThreadRightComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MessengerReplyThreadRightComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
