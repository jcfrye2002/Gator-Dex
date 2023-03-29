import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateDexComponent } from './create-dex.component';

describe('CreateDexComponent', () => {
  let component: CreateDexComponent;
  let fixture: ComponentFixture<CreateDexComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateDexComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateDexComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
