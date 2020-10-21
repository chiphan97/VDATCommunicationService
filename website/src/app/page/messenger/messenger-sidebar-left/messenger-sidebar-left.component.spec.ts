import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MessengerSidebarLeftComponent } from './messenger-sidebar-left.component';

describe('MessengerSidebarLeftComponent', () => {
  let component: MessengerSidebarLeftComponent;
  let fixture: ComponentFixture<MessengerSidebarLeftComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MessengerSidebarLeftComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MessengerSidebarLeftComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
