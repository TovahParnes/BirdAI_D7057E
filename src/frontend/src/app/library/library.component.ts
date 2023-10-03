import {Component} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';

@Component({
  selector: 'app-library',
  templateUrl: './library.component.html',
  styleUrls: ['./library.component.css']
})

export class LibraryComponent {

  constructor(
    private router: Router, 
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService) {
  }

  logout(): void {
    this.socialAuthService.signOut().then(() => this.router.navigate(['login']));
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

  toggleTheme(): void {
    //this.mainApp.switchDarkmodeSetting();
  }

  navigateToSpecies(): void {
    this.router.navigate(['species-page']);
  }
}
