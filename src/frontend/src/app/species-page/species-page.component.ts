import {Component, ElementRef, AfterViewInit, OnInit} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider} from '@abacritt/angularx-social-login';
import {Router, ActivatedRoute} from '@angular/router';
import { AppComponent } from '../app.component';
import { HttpClient } from '@angular/common/http';
import { Location } from '@angular/common';

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
  imageDate!: string;
  imageDesc!: string;
  responseData: ApiResponse | null = null;
  images: string[] = [];

  constructor(
    private router: Router,
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private elRef: ElementRef,
    private http: HttpClient,
    private route: ActivatedRoute,
    private location: Location,
    ) {

  }

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.imageId = decodeURIComponent(params['imageId']);
      this.imageName = decodeURIComponent(params['imageName'])
      this.imageDate = decodeURIComponent(params['imageDate'])
      this.imageDesc = decodeURIComponent(params['imageDesc'])
        if (this.imageDate == "undefined"){
          this.imageDate = "Not Found Yet"
        }
      

      this.images = [
        this.imageId,
      ];
    });
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
    const imageWidth = imageElement.width;
    const buttonContainerWidth = buttonContainer.clientWidth;
    const leftPosition = (imageWidth - buttonContainerWidth) / 2 + 'px';
    buttonContainer.style.left = leftPosition;
  }

  goBack(): void {
    this.location.back();
  }

}
