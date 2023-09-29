import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {CommonModule} from "@angular/common";

import {AppComponent} from './app.component';
import {RouterModule} from '@angular/router';
import {LoginComponent} from './login/login.component';
import {MainPageComponent} from './home/home-page.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
//import {MatCardModule} from '@angular/material/card';
//import {MatFormFieldModule} from '@angular/material/form-field';
//import {MatButtonModule} from '@angular/material/button';
//import {MatInputModule} from '@angular/material/input';
import {GoogleLoginProvider, SocialLoginModule} from 'angularx-social-login';
import {AuthGuardService} from './auth-guard.service';
import { LibraryComponent } from './library/library.component';
import { SpeciesPageComponent } from './species-page/species-page.component';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    MainPageComponent,
    LibraryComponent,
    SpeciesPageComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot([
      {path: 'login', component: LoginComponent},
      {path: 'mainpage', component: MainPageComponent, canActivate: [AuthGuardService]},
      {path: 'library', component: LibraryComponent},
      {path: 'species-page', component: SpeciesPageComponent},
      {path: '**', component: LoginComponent},
    ]),
    BrowserAnimationsModule,
    CommonModule,
    //MatCardModule,
    //MatFormFieldModule,
    //MatButtonModule,
    //MatInputModule,
    SocialLoginModule,

  ],
  exports:[RouterModule],
  providers: [{
    provide: 'SocialAuthServiceConfig',
    useValue: {
      autoLogin: true,
      providers: [
        {
          id: GoogleLoginProvider.PROVIDER_ID,
          provider: new GoogleLoginProvider('148517665605-jspahbqleats6lvlag9kasc2c11b5g7o.apps.googleusercontent.com')
        }
      ]
    }
  },
    AuthGuardService],
  bootstrap: [AppComponent]
})
export class AppModule {
}