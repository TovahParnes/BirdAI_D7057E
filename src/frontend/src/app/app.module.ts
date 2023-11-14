// base imports
import {NgModule} from '@angular/core';
import {RouterModule} from '@angular/router';
import {BrowserModule} from '@angular/platform-browser';
import {CommonModule} from "@angular/common";
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {HttpClientModule} from '@angular/common/http';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

// page components imports
import {AppComponent} from './app.component';
import {AboutComponent} from './about/about.component';
import {LoginComponent} from './login/login.component';
import {MainPageComponent} from './home/home-page.component';
import {LibraryComponent} from './library/library.component';
import {TakenImagesPageComponent} from './taken-images-page/taken-images-page.component';
import {SpeciesPageComponent} from './species-page/species-page.component';
import {ProfilePageComponent} from './profile-page/profile-page.component';
import {NavbarComponent} from './navbar/navbar.component';
import {CardComponent, BottomInfoComponent} from './card/card.component';
import {FirstPageComponent} from './first-page/first-page.component';
import { AdminComponent } from './admin/admin.component';

// login authguard imports
import {GoogleLoginProvider, GoogleSigninButtonModule, SocialLoginModule, SocialAuthServiceConfig} from '@abacritt/angularx-social-login';
import {AuthGuardService} from './services/auth-guard.service';
import {AuthGuardAdminService} from './services/auth-admin-guard.service'

// material
import {MatButtonModule} from '@angular/material/button';
import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import {MatStepperModule} from '@angular/material/stepper';
import {MatRadioModule} from '@angular/material/radio';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatCardModule} from '@angular/material/card';
import {MatIconModule} from '@angular/material/icon';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {MatButtonToggleModule} from '@angular/material/button-toggle';





@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    MainPageComponent,
    LibraryComponent,
    SpeciesPageComponent,
    TakenImagesPageComponent,
    ProfilePageComponent,
    CardComponent,
    NavbarComponent,
    BottomInfoComponent,
    FirstPageComponent,
    AboutComponent,
    AdminComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot([
      {path: 'login', component: LoginComponent},
      {path: 'mainpage', component: MainPageComponent},
      {path: 'library', component: LibraryComponent},
      {path: 'takenImages', component: TakenImagesPageComponent},
      {path: 'profile', component: ProfilePageComponent}, //canActivate: [AuthGuardService]
      {path: 'species-page', component: SpeciesPageComponent},
      {path: 'first-page', component: FirstPageComponent},
      {path: 'about', component: AboutComponent},
      {path: 'admin',component: AdminComponent, canActivate: [AuthGuardAdminService]},
      {path: '**', component: FirstPageComponent},
    ]),
    BrowserAnimationsModule,
    CommonModule,
    GoogleSigninButtonModule,
    SocialLoginModule,
    FormsModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatSlideToggleModule,
    MatStepperModule,
    MatRadioModule,
    MatFormFieldModule,
    MatCardModule,
    MatIconModule,
    MatProgressSpinnerModule,
    MatButtonToggleModule,
  ],
  exports:[
    RouterModule
  ],
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
    AuthGuardAdminService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}