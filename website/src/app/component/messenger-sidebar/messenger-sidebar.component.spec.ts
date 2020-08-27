import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MessengerSidebarComponent } from './messenger-sidebar.component';

describe('MessengerSidebarComponent', () => {
  let component: MessengerSidebarComponent;
  let fixture: ComponentFixture<MessengerSidebarComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MessengerSidebarComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MessengerSidebarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
