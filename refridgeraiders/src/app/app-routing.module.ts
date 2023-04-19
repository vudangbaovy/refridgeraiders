import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from './home/home.component';
import { CommonModule } from '@angular/common';
import { RegisterComponent } from './register/register.component';
import { ProfileComponent } from './profile/profile.component';
import { RecipeComponent } from './recipe/recipe.component';

const routes: Routes = [ 
    { path: '', component: HomeComponent },
    { path: 'login', component: LoginComponent },
    { path: 'register', component: RegisterComponent },
    { path: 'profile', component: ProfileComponent },
    { path: 'recipe', component: RecipeComponent }

];

@NgModule({
    imports: [
        CommonModule,
        RouterModule.forRoot(routes)
    ],
    exports: [RouterModule]
})
export class AppRoutingModule { }
