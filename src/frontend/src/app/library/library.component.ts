import {Component, ViewChildren, QueryList} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';
import { CardComponent } from '../card/card.component';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {getAllBirdsResponse } from 'src/assets/components/components';
import { environment } from 'src/environments/environment';
import { WikiPageSegment, WikiSummary, WikirestService } from '../services/wiki.service';

@Component({
  selector: 'app-library',
  templateUrl: './library.component.html',
  styleUrls: ['./library.component.css'],
})

export class LibraryComponent {

  private jsonUrl = 'assets/data.json';
  cardlist: any[] = [];
  foundlist: any[] = [];
  color = "#2196f3"
  private lastLetter = ""
  private found = false
  alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split('');
  @ViewChildren(CardComponent)
  cards!: QueryList<CardComponent>;

  // cardlist = [
  //   {title: 'Duck',imageSrc: 'assets/duck.jpg'},
  //   {title: 'Kangaroo',imageSrc: 'assets/duck.jpg'},
  //   {title: 'Elephant',imageSrc: 'assets/undulat.jpg'},
  //   {title: 'Duck',imageSrc: 'assets/duck.jpg'},
  //   {title: 'Kangaroo',imageSrc: 'assets/duck.jpg'},
  //   {title: 'Elephant',imageSrc: 'assets/undulat.jpg'},
  // ]

  // foundlist = [
  //   {title: 'Duck',imageSrc: 'assets/duck.jpg'},
  //   {title: 'Kangaroo',imageSrc: 'assets/duck.jpg'},
  // ]

  private backupCards: any[] = []

  constructor(
    private router: Router,
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private http: HttpClient,
    private wikiRest: WikirestService,
    ) {
  }

  navigateToSpecies(imageId: string, imageName: string,imageSound:string, imageDesc: string, imageGenus:Boolean): void {
    this.router.navigate(['species-page'], {
      queryParams: {
        imageId: encodeURIComponent(imageId),
        imageName: encodeURIComponent(imageName),
        imageSound: encodeURIComponent(imageSound),
        imageDesc: encodeURI(imageDesc),
        imageGenus: imageGenus
      }
      });
  }

  allBirds: getAllBirdsResponse = {
    data:[],
    timestamp: ""
  }

  allBirdsBackup: getAllBirdsResponse = {
    data:[],
    timestamp: ""
  }

  //sorts by name but should prob not be visible as a button but rather done autmatically in the background
  sortCards() {
    this.cardlist.sort((a, b) => a.title.localeCompare(b.title));
    this.allBirds.data.sort((a, b) => a.Name.localeCompare(b.Name));
  }

  // filterCards(letter: string): void {
  //   if(this.found == true){
  //     this.allBirds.data = this.foundlist
  //   }else{
  //     this.allBirds.data = this.backupCards
  //   }
  //   this.allBirds.data = this.allBirds.data.filter(card => card.Name.startsWith(letter));
  //   this.lastLetter = letter
  // }

  filterCards(letter: string): void {
    this.allBirds.data = this.allBirdsBackup.data;
    this.allBirds.data = this.allBirds.data.filter(card => card.Name.startsWith(letter));
    this.lastLetter = letter;
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

  getData(): Observable<any[]> {
    return this.http.get<any[]>(this.jsonUrl);
  }

  ngOnInit(): void {
    this.getAllBirds();
    this.cardlist = this.allBirds.data;
    this.foundlist = this.allBirds.data;
  }

  getAllBirds(){
    this.sendGetBirds().subscribe(
      (response: getAllBirdsResponse) => {
        this.allBirds = response;
        this.allBirdsBackup.data = response.data;
        for(let i = 0; i <= this.allBirds.data.length; i++){
          this.setDataImageToWikiImage(this.getWikiLinkTitle(i),i);
        }
      },
      err => { 
        console.error("Failed at sending data:" + err); 
      }
    );
  }
  sendGetBirds() {
    return this.http.get<getAllBirdsResponse>(environment.identifyRequestURL+"/birds/list?set=0");
  }

  getWikiLinkTitle(index:number){
    let cutOffIndex = this.allBirds.data[index].Description.indexOf('wiki/');
    let cutString = this.allBirds.data[index].Description.substring(cutOffIndex + 'wiki/'.length)
    return cutString;
  }

  async setDataImageToWikiImage(wikiTitle:string,index:number){
    this.wikiRest.getWiki(wikiTitle).subscribe(data => {
      console.log(data);
      if(data.extract){
      //this.wikiData = data;
      if(data.originalimage?.source){
        this.allBirds.data[index].Image.data = data.originalimage?.source;
        this.allBirdsBackup.data[index].Image.data = data.originalimage?.source;
      }
      }
    }, err => { console.log('something went wrong' + err)
  }); 

}
}
