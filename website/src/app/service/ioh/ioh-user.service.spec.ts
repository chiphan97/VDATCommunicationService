import { TestBed } from '@angular/core/testing';

import { IohUserService } from './ioh-user.service';

describe('IohUserService', () => {
  let service: IohUserService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(IohUserService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
