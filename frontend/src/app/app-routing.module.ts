import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

//Local components
import { UserListComponent } from './user-list/user-list.component';
import { UpdateUserComponent } from './update-user/update-user.component';
import { LoginComponent } from './login/login.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { UserDetailsComponent } from './user-details/user-details.component';
import { CreateUserComponent } from './create-user/create-user.component';

const routes: Routes = [
   { path: 'login', component:LoginComponent},
   { path: 'new-user', component: CreateUserComponent},
   { path: 'update-user', component: UpdateUserComponent},
   { path: 'user-details', component: UserDetailsComponent},
   { path: 'user-list', component: UserListComponent},
   { path: '', redirectTo: 'login', pathMatch: 'full'},
   {path: '**', component:PageNotFoundComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }