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

  color = "#2196f3"
  private lastLetter = ""
  private found = false
  alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split('');
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

  foundlist = [
    {title: 'Duck',imageSrc: 'assets/duck.jpg'},
    {title: 'Kangaroo',imageSrc: 'assets/duck.jpg'},
  ]

  private backupCards = this.cardlist;

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

  //sorts by name but should prob not be visible as a button but rather done autmatically in the background
  sortCards() {
    this.cardlist.sort((a, b) => a.title.localeCompare(b.title));
  }

  filterCards(letter: string): void {
    if(this.found == true){
      this.cardlist = this.foundlist
    }else{
      this.cardlist = this.backupCards
    }
    this.cardlist = this.cardlist.filter(card => card.title.startsWith(letter));
    this.lastLetter = letter
  }

  toggleFound(){
    if (this.found == true){
      this.found = false
      this.cardlist = this.backupCards
      this.filterCards(this.lastLetter)
      this.color = "#2196f3"
    }else{
      this.found = true
      this.cardlist = this.foundlist
      this.filterCards(this.lastLetter)
      this.color = "#1971B8"
    }
  }
}
