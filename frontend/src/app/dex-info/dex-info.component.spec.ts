import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DexInfoComponent } from './dex-info.component';

describe('DexInfoComponent', () => {
  let component: DexInfoComponent;
  let fixture: ComponentFixture<DexInfoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DexInfoComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DexInfoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
