import {Component, ElementRef} from '@angular/core';
import {SocialAuthService} from '@abacritt/angularx-social-login';
import {Router, ActivatedRoute} from '@angular/router';
import {AppComponent} from '../app.component';
import {HttpClient} from '@angular/common/http';
import {Location} from '@angular/common';
import {WikiImages, WikiPageSegment, WikiSummary, WikirestService} from '../services/wiki.service';
import {ApiResponse} from 'src/assets/components/components';

@Component({
  selector: 'app-species-page',
  templateUrl: './species-page.component.html',
  styleUrls: ['./species-page.component.css']
})

export class SpeciesPageComponent{
  imageId!: string;
  imageName!: string;
  imageDate!: string;
  imageDesc!: string;
  imageSound!: string;
  imageGenus!: Boolean;
  imagePage!: number;
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
    window.scrollTo(0, 0);
    this.route.queryParams.subscribe(params => {
      this.imageId = decodeURIComponent(params['imageId']);
      this.imageName = decodeURIComponent(params['imageName']);
      this.imageSound = decodeURIComponent(params['imageSound']);
      this.imageDesc = decodeURIComponent(params['imageDesc']);
      this.imageGenus = params['imageGenus'] === 'true';
      this.imagePage = params['imagePage'];
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

  get currentImage(): string {
    return this.images[0];
  }

  goBack(): void {
    if(this.imagePage!= null){
      this.router.navigate(['library'], {
        queryParams: {
          imagePage: this.imagePage
        }
    });
  }else{
    this.location.back();
  }
}

  getWikiLinkTitle(){
    let cutOffIndex = this.imageDesc.indexOf('wiki/');
    let cutString = this.imageDesc.substring(cutOffIndex + 'wiki/'.length)
    return cutString;
  }

  getWikiData(wikiTitle:string){
    this.wikiRest.getWiki(wikiTitle).subscribe(data => {
      if(data.extract){
      this.wikiData = data;
      }
    }, err => { console.log('something went wrong' + err)
  }); 
  }

  getWikiImage(wikiTitle:string){
    this.wikiRest.getWikiImages(wikiTitle).subscribe((data: WikiImages) => {
      for(let i=0;i<= data.items.length;i++){
        if(data.items[i].title.includes('map')){
          this.wikiImages = data.items[i].srcset[0].src;
        }
      }
    })
  }

}
