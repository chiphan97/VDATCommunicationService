import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MessengerMessageComponent } from './messenger-message.component';

describe('MessengerMessageComponent', () => {
  let component: MessengerMessageComponent;
  let fixture: ComponentFixture<MessengerMessageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MessengerMessageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MessengerMessageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
