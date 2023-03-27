import { UserDetailsComponent } 
   from './user-details/user-details.component';
import { CreateUserComponent } 
   from './create-user/create-user.component';
import { NgModule } from '@angular/core';
import { Routes, RouterModule } 
   from '@angular/router';
import {UserListComponent } 
  from './user-list/user-list.component';
import { UpdateUserComponent } 
   from './update-user/update-user.component';
import { LoginComponent } from './login/login.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';

const routes: Routes = [
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  { path: 'user-list', component: UserListComponent },
  { path: 'new-user', component: CreateUserComponent },
  { path: 'update/:id', component: UpdateUserComponent },
  { path: 'details/:id', component: UserDetailsComponent },
  {path: 'login', component:LoginComponent},
  {path: '**', component:PageNotFoundComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }