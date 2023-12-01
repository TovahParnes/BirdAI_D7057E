import {Component, OnInit} from '@angular/core';
import {SocialAuthService} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';
import { CardComponent } from '../card/card.component';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {getAllBirdsResponse } from 'src/assets/components/components';
import { environment } from 'src/environments/environment';
import {FormControl} from '@angular/forms';
import { WikiPageSegment, WikiSummary, WikirestService } from '../services/wiki.service';

@Component({
  selector: 'app-library',
  templateUrl: './library.component.html',
  styleUrls: ['./library.component.css'],
})

export class LibraryComponent implements OnInit {

  jsonUrl = 'assets/data.json';
  alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split('');
  searchInput: FormControl = new FormControl();
  pageSearch: FormControl = new FormControl();
  selectedOption: FormControl = new FormControl('');
  cardlist: any[] = [];
  foundlist: any[] = [];
  currentPageNumber: Number = 0;
  
  constructor(
    private router: Router,
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private http: HttpClient,
    private wikiRest: WikirestService,
    ) {
  }

  ngOnInit() {
    this.getAllBirds(0);
    this.cardlist = this.allBirds.data;
    this.foundlist = this.allBirds.data;

    this.selectedOption.valueChanges.subscribe(value => {
      this.filterByLetter(value);
    });

    this.searchInput.valueChanges.subscribe(value => {
      this.filterBySearch(value.toUpperCase());
    });

    this.pageSearch.valueChanges.subscribe(value => {
      const numericValue = parseInt(value, 10);
      this.currentPageNumber = numericValue;
      if (Number.isNaN(this.currentPageNumber.valueOf())){
        this.currentPageNumber = 0;
      }
      this.changePage(0);
    });
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

  sortCards() {
    this.cardlist.sort((a, b) => a.title.localeCompare(b.title));
    this.allBirds.data.sort((a, b) => a.Name.localeCompare(b.Name));
  }

  filterByLetter(selectedValue: any) {
    this.allBirds.data = this.allBirdsBackup.data;
    this.allBirds.data = this.allBirds.data.filter(card => card.Name.startsWith(selectedValue));
  }

  filterBySearch(searchValue: string) {
    this.allBirds.data = this.allBirdsBackup.data;
    this.allBirds.data = this.allBirds.data.filter(card => card.Name.match(searchValue));
  }

  getData(): Observable<any[]> {
    return this.http.get<any[]>(this.jsonUrl);
  }

  getAllBirds(pageNumber:Number){
    this.sendRequestGetBirds(pageNumber).subscribe(
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

  changePage(increment:Number){
    if(increment.valueOf()<0){
    this.currentPageNumber = this.currentPageNumber.valueOf()-1;
    }else if(increment.valueOf()>0){
      this.currentPageNumber = this.currentPageNumber.valueOf()+1;
    }
    if (this.currentPageNumber.valueOf() < 0){
      this.currentPageNumber = 0;
    }else{
      this.getAllBirds(this.currentPageNumber);
    }
  }

  sendRequestGetBirds(pageNumber:Number) {
    return this.http.get<getAllBirdsResponse>(environment.identifyRequestURL+"/birds/list?set="+pageNumber);
  }

  getWikiLinkTitle(index:number){
    let cutOffIndex = this.allBirds.data[index].Description.indexOf('wiki/');
    let cutString = this.allBirds.data[index].Description.substring(cutOffIndex + 'wiki/'.length)
    return cutString;
  }

  async setDataImageToWikiImage(wikiTitle:string,index:number){
    this.wikiRest.getWiki(wikiTitle).subscribe(data => {
      if(data.extract){
      if(data.originalimage?.source){
        this.allBirds.data[index].Image = data.originalimage?.source;
        this.allBirdsBackup.data[index].Image = data.originalimage?.source;
      }
      }
    }, err => { console.log('something went wrong' + err)
  }); 

}
}
