import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TimeslotsComponent } from './timeslots.component';

describe('TimeslotsComponent', () => {
  let component: TimeslotsComponent;
  let fixture: ComponentFixture<TimeslotsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TimeslotsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TimeslotsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
