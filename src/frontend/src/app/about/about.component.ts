import { Component } from '@angular/core';
import {SocialAuthService, GoogleLoginProvider} from '@abacritt/angularx-social-login';
import { Router } from '@angular/router';
import { AppComponent } from '../app.component';

@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.css']
})
export class AboutComponent {

  constructor(
    private router: Router,
    public socialAuthService: SocialAuthService,
    public mainApp: AppComponent,
  ){

  }

  navigateToHome(): void {
    this.router.navigate(['mainpage']);
  }

  navigateToTakenImages(): void {
    this.router.navigate(['takenImages']);
  }

  navigateToLibrary(): void {
    this.router.navigate(['library']);
  }

  navigateToProfilePage(): void {
    this.router.navigate(['profile']);
  }

  logout(): void {
    this.socialAuthService.signOut().then(() => this.router.navigate(['login']));
  }
}
