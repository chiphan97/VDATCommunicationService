import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ListUserChatComponent } from './list-user-chat.component';

describe('ListUserChatComponent', () => {
  let component: ListUserChatComponent;
  let fixture: ComponentFixture<ListUserChatComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ListUserChatComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ListUserChatComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
