import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewDexComponent } from './view-dex.component';

describe('ViewDexComponent', () => {
  let component: ViewDexComponent;
  let fixture: ComponentFixture<ViewDexComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewDexComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewDexComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
