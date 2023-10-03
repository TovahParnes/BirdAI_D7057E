// base imports
import {NgModule} from '@angular/core';
import {RouterModule} from '@angular/router';
import {BrowserModule} from '@angular/platform-browser';
import {CommonModule} from "@angular/common";
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule} from '@angular/forms';

// page components imports
import {AppComponent} from './app.component';
import {LoginComponent} from './login/login.component';
import {MainPageComponent} from './home/home-page.component';
import {LibraryComponent} from './library/library.component';
import {TakenImagesPageComponent} from './taken-images-page/taken-images-page.component';
import {SpeciesPageComponent} from './species-page/species-page.component';
import {ProfilePageComponent} from './profile-page/profile-page.component';

// login authguard imports
import {GoogleLoginProvider, GoogleSigninButtonDirective, GoogleSigninButtonModule, SocialLoginModule, SocialAuthServiceConfig} from '@abacritt/angularx-social-login';
import {AuthGuardService} from './services/auth-guard.service';

// material
import {MatButtonModule} from '@angular/material/button';
import {MatSlideToggleModule} from '@angular/material/slide-toggle';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    MainPageComponent,
    LibraryComponent,
    SpeciesPageComponent,
    TakenImagesPageComponent,
    ProfilePageComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot([
      {path: 'login', component: LoginComponent},
      {path: 'mainpage', component: MainPageComponent, canActivate: [AuthGuardService]},
      {path: 'library', component: LibraryComponent, canActivate: [AuthGuardService]},
      {path: 'takenImages', component: TakenImagesPageComponent, canActivate: [AuthGuardService]},
      {path: 'profile', component: ProfilePageComponent, canActivate: [AuthGuardService]},
      {path: 'species-page', component: SpeciesPageComponent, canActivate: [AuthGuardService]},
      {path: '**', component: LoginComponent},
    ]),
    BrowserAnimationsModule,
    CommonModule,
    GoogleSigninButtonModule,
    SocialLoginModule,
    MatButtonModule,
    MatSlideToggleModule,
    FormsModule,
  ],
  exports:[RouterModule],
  providers: [
    {
      provide: 'SocialAuthServiceConfig',
      useValue: {
        autoLogin: false,
        providers: [
          {
            id: GoogleLoginProvider.PROVIDER_ID,
            provider: new GoogleLoginProvider('1048676205865-lm62p2961pftdfrj919m4fts8lel1hu9.apps.googleusercontent.com')
          },
        ],
        onError: (err) => {
          console.error(err);
        }
      } as SocialAuthServiceConfig,
    },
    AuthGuardService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}