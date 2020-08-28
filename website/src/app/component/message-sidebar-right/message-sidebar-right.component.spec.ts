import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MessageSidebarRightComponent } from './message-sidebar-right.component';

describe('MessageSidebarRightComponent', () => {
  let component: MessageSidebarRightComponent;
  let fixture: ComponentFixture<MessageSidebarRightComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MessageSidebarRightComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MessageSidebarRightComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
