import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { IntegratedComponent } from './integrated.component';

describe('IntegratedComponent', () => {
  let component: IntegratedComponent;
  let fixture: ComponentFixture<IntegratedComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ IntegratedComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(IntegratedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
