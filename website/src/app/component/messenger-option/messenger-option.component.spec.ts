import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MessengerOptionComponent } from './messenger-option.component';

describe('MessengerOptionComponent', () => {
  let component: MessengerOptionComponent;
  let fixture: ComponentFixture<MessengerOptionComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MessengerOptionComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MessengerOptionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
