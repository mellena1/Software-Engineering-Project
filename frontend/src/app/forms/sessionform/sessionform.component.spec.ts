import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SessionformComponent } from './sessionform.component';

describe('SessionformComponent', () => {
  let component: SessionformComponent;
  let fixture: ComponentFixture<SessionformComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SessionformComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SessionformComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
