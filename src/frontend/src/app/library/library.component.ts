import {Component} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';
import { Card2Component } from '../card/card.component';

@Component({
  selector: 'app-library',
  templateUrl: './library.component.html',
  styleUrls: ['./library.component.css']
})

export class LibraryComponent {

  cardlist = [
    {title: 'Duck',imageSrc: 'assets/duck.jpg', date:'2023-10-05'},
    {title: 'Budgie',imageSrc: 'assets/undulat.jpg', date:'2023-10-04'},
  ]

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

  navigateToSpecies(imageId: string, imageName: string): void {
    this.router.navigate(['species-page'], {
      queryParams: {
        imageId: encodeURIComponent(imageId),
        imageName: encodeURIComponent(imageName),
      }
      });
  }
}
