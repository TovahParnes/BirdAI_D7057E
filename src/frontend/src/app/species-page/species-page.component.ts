import {Component, ElementRef, AfterViewInit, OnInit} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider} from '@abacritt/angularx-social-login';
import {Router, ActivatedRoute} from '@angular/router';
import { AppComponent } from '../app.component';
import { HttpClient } from '@angular/common/http';

interface ApiResponse {
  data: {  
    id : string;
    authId: string;
    createdAt: string;
    username: string;
  }[];
  message: string;
  success: boolean;
}

@Component({
  selector: 'app-species-page',
  templateUrl: './species-page.component.html',
  styleUrls: ['./species-page.component.css']
})

export class SpeciesPageComponent implements AfterViewInit{
  imageId!: string;
  imageName!: string;
  responseData: ApiResponse | null = null;
  images: string[] = [];

  constructor(
    private router: Router,
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private elRef: ElementRef,
    private http: HttpClient,
    private route: ActivatedRoute,
    ) {

  }

  ngOnInit() {
    // Make an HTTP GET request to the Swagger service's API
    this.http.get<ApiResponse>('http://localhost:4000/swagger/index.html').subscribe(data => {
      this.responseData = data;
      console.log(data);
    });
    this.route.queryParams.subscribe(params => {
      this.imageId = decodeURIComponent(params['imageId']);
      this.imageName = decodeURIComponent(params['imageName'])

      this.images = [
        this.imageId,
        'assets/map.png',
      ];
    });
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

  toggleTheme(): void {
    //this.mainApp.switchDarkmodeSetting();
  }

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
