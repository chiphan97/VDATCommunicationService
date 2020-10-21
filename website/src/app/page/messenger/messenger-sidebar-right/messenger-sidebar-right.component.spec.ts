import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MessengerSidebarRightComponent } from './messenger-sidebar-right.component';

describe('MessengerSidebarRightComponent', () => {
  let component: MessengerSidebarRightComponent;
  let fixture: ComponentFixture<MessengerSidebarRightComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MessengerSidebarRightComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MessengerSidebarRightComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
