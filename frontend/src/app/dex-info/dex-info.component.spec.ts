import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { DexInfoComponent } from './dex-info.component';
import { MatFormFieldModule } from '@angular/material/form-field';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { MatInputModule } from '@angular/material/input';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';



describe('DexInfoComponent', () => {
  let component: DexInfoComponent;
  let fixture: ComponentFixture<DexInfoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, MatFormFieldModule, 
        BrowserModule,
        FormsModule,
        ReactiveFormsModule,
      MatInputModule,
      BrowserAnimationsModule],
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
