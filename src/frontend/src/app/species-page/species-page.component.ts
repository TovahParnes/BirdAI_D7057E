import {Component, ElementRef, AfterViewInit} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';

@Component({
  selector: 'app-species-page',
  templateUrl: './species-page.component.html',
  styleUrls: ['./species-page.component.css']
})

export class SpeciesPageComponent implements AfterViewInit{

  constructor(
    private router: Router,
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private elRef: ElementRef
    ) {

  }

  moveToHome(){
    this.router.navigate(['mainpage']);
  }

  moveToLibrary(): void {
    this.router.navigate(['library']);
  }

  toggleTheme(): void {
    //this.mainApp.switchDarkmodeSetting();
  }

  images: string[] = [
    'assets/duck.jpg',
    'assets/undulat.jpg',
    'assets/map.png',
  ];
  currentImageIndex = 0;

  ngAfterViewInit() {
    this.updateButtonPosition();
  }

  get currentImage(): string {
    return this.images[this.currentImageIndex];
  }

  nextImage() {
    this.currentImageIndex = (this.currentImageIndex + 1) % this.images.length;
    this.updateButtonPosition();
  }

  prevImage() {
    this.currentImageIndex = (this.currentImageIndex - 1 + this.images.length) % this.images.length;
    this.updateButtonPosition();
  }

  updateButtonPosition() {
    const imageElement = this.elRef.nativeElement.querySelector('#slideshow-image');
    const buttonContainer = this.elRef.nativeElement.querySelector('.button-container');

    // Adjust the left position of the button container based on the image width
    const imageWidth = imageElement.width;
    const buttonContainerWidth = buttonContainer.clientWidth;
    const leftPosition = (imageWidth - buttonContainerWidth) / 2 + 'px';

    buttonContainer.style.left = leftPosition;
  }

}
