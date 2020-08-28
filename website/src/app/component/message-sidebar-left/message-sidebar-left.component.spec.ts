import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MessageSidebarLeftComponent } from './message-sidebar-left.component';

describe('MessageSidebarLeftComponent', () => {
  let component: MessageSidebarLeftComponent;
  let fixture: ComponentFixture<MessageSidebarLeftComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MessageSidebarLeftComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MessageSidebarLeftComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
