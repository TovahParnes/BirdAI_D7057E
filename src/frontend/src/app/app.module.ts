import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

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
import { UploadImageComponent } from './upload-image/upload-image.component';


@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    MainPageComponent,
    UploadImageComponent,
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot([
      {path: 'login', component: LoginComponent},
      {path: 'mainpage', component: MainPageComponent, canActivate: [AuthGuardService]},
      {path: '**', component: LoginComponent},
      {path: 'upload-page', component: UploadImageComponent}
    ]),
    BrowserAnimationsModule,
    //MatCardModule,
    //MatFormFieldModule,
    //MatButtonModule,
    //MatInputModule,
    SocialLoginModule
  ],
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