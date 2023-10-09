import {Component, ViewChildren, QueryList} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';
import { CardComponent } from '../card/card.component';

@Component({
  selector: 'app-taken-images-page',
  templateUrl: './taken-images-page.component.html',
  styleUrls: ['./taken-images-page.component.css'],
})

export class TakenImagesPageComponent {

  @ViewChildren(CardComponent)
  cards!: QueryList<CardComponent>;

  cardlist = [
    {title: 'Duck',imageSrc: 'assets/duck.jpg'},
    {title: 'Kangaroo',imageSrc: 'assets/duck.jpg'},
    {title: 'Elephant',imageSrc: 'assets/undulat.jpg'},
    {title: 'Duck',imageSrc: 'assets/duck.jpg'},
    {title: 'Kangaroo',imageSrc: 'assets/duck.jpg'},
    {title: 'Elephant',imageSrc: 'assets/undulat.jpg'},
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

  sortCards() {
    // const sortedCards = this.cards.toArray().sort((a, b) => a.title.localeCompare(b.title));
    // this.cards.reset(sortedCards);
    this.cardlist.sort((a, b) => a.title.localeCompare(b.title));
  }
}
