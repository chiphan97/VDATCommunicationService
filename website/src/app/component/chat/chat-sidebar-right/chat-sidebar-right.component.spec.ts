import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ChatSidebarRightComponent } from './chat-sidebar-right.component';

describe('ChatSidebarRightComponent', () => {
  let component: ChatSidebarRightComponent;
  let fixture: ComponentFixture<ChatSidebarRightComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ChatSidebarRightComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChatSidebarRightComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
