import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

// used to create fake backend
import { fakeBackendProvider } from './_helpers';
import { AlertModule } from './_alert'

import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';

import { JwtInterceptor, ErrorInterceptor } from './_helpers';
import { LoginComponent } from './login';;
import { RegisterComponent } from './register/register.component'
import { environment } from '@environments/environment';;
import { GamesComponent } from './games/games.component'

function providers(){
  let providers: any = [
    { provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true },
    { provide: HTTP_INTERCEPTORS, useClass: ErrorInterceptor, multi: true },
  ]
  if (environment.noBack){
    providers.push(fakeBackendProvider)
  }
  return providers;
}

@NgModule({
    imports: [
        BrowserModule,
        ReactiveFormsModule,
        HttpClientModule,
        AppRoutingModule,
        AlertModule
    ],
    declarations: [
        AppComponent,
        GamesComponent,
        LoginComponent
,
        RegisterComponent ,
        GamesComponent   ],
    providers: providers(),
    bootstrap: [AppComponent]
})

export class AppModule { }
