import { BrowserModule } from '@angular/platform-browser';
import { Component, NgModule } from '@angular/core';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CreateUserComponent } from './create-user/create-user.component';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule} from '@angular/platform-browser/animations'

//Forms
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

//Material
import { MatFormFieldModule } from '@angular/material/form-field'
import { MatInputModule } from '@angular/material/input'
import { MatToolbarModule } from '@angular/material/toolbar'

//local shit
import { LoginComponent } from './login/login.component';
import { UpdateUserComponent } from './update-user/update-user.component';
import { UserDetailsComponent } from './user-details/user-details.component';
import { UserListComponent } from './user-list/user-list.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';

@NgModule({
  declarations: [
    AppComponent,
    CreateUserComponent,
    UserDetailsComponent,
    UserListComponent,
    UpdateUserComponent,
    LoginComponent,
    PageNotFoundComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    MatFormFieldModule,
    ReactiveFormsModule,
    MatInputModule,
    BrowserAnimationsModule,
    MatToolbarModule,
  ],
  exports:[
    FormsModule,
    MatFormFieldModule,
    MatInputModule,
    ReactiveFormsModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

export class MaterialModule{ }
