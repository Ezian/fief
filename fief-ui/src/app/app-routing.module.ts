import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { GamesComponent } from './games/games.component';


import { LoginComponent } from './login';
import { RegisterComponent } from './register/register.component';
import { AuthGuard } from './_helpers';

const routes: Routes = [
    { path: '', component: GamesComponent, canActivate: [AuthGuard] },
    { path: 'login', component: LoginComponent },
    { path: 'signup', component: RegisterComponent },

    // otherwise redirect to games
    { path: '**', redirectTo: '' }
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule { }
