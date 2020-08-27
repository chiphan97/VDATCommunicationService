import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MessengerDrawerComponent } from './messenger-drawer.component';

describe('MessengerDrawerComponent', () => {
  let component: MessengerDrawerComponent;
  let fixture: ComponentFixture<MessengerDrawerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MessengerDrawerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MessengerDrawerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
