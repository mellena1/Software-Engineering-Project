import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SpeakerformComponent } from './speakerform.component';

describe('SpeakerformComponent', () => {
  let component: SpeakerformComponent;
  let fixture: ComponentFixture<SpeakerformComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SpeakerformComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SpeakerformComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
