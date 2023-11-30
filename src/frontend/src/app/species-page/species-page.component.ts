import {Component, ElementRef, AfterViewInit, OnInit} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider} from '@abacritt/angularx-social-login';
import {Router, ActivatedRoute} from '@angular/router';
import { AppComponent } from '../app.component';
import { HttpClient } from '@angular/common/http';
import { Location } from '@angular/common';
import { WikiImages, WikiPageSegment, WikiSummary, WikirestService } from '../services/wiki.service';

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
  imageSound!: string;
  imageGenus!: Boolean;
  responseData: ApiResponse | null = null;
  images: string[] = [];
  wikiData: WikiSummary = new WikiSummary;
  wikiContent: WikiPageSegment = new WikiPageSegment;
  wikiImages: string = "";

  constructor(
    private router: Router,
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private elRef: ElementRef,
    private http: HttpClient,
    private route: ActivatedRoute,
    private location: Location,
    private wikiRest: WikirestService,
    ) {

  }

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.imageId = decodeURIComponent(params['imageId']);
      this.imageName = decodeURIComponent(params['imageName']);
      this.imageSound = decodeURIComponent(params['imageSound']);
      this.imageDesc = decodeURIComponent(params['imageDesc']);
      this.imageGenus = params['imageGenus'];
        if (this.imageDate == "undefined"){
          this.imageDate = "Not Found Yet"
        }
      

      this.images = [
        this.imageId,
      ];
    });
    this.getWikiData(this.getWikiLinkTitle());
    this.getWikiImage(this.getWikiLinkTitle());
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

  getWikiLinkTitle(){
    let cutOffIndex = this.imageDesc.indexOf('wiki/');
    let cutString = this.imageDesc.substring(cutOffIndex + 'wiki/'.length)
    return cutString;
  }

  getWikiData(wikiTitle:string){
    this.wikiRest.getWiki(wikiTitle).subscribe(data => {
      console.log(data);
      if(data.extract){
      this.wikiData = data;
      }
    }, err => { console.log('something went wrong' + err)
  }); 
  
  }

  getWikiImage(wikiTitle:string){
    this.wikiRest.getWikiImages(wikiTitle).subscribe((data: WikiImages) => {
      console.log(data);
      for(let i=0;i<= data.items.length;i++){
        if(data.items[i].title.includes('map')){
          this.wikiImages = data.items[i].srcset[0].src;
        }
      }
    })
  }

}
