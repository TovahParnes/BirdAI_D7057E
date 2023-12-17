import {Component, OnInit, ViewChild, OnDestroy} from '@angular/core';
import {Router, ActivatedRoute} from '@angular/router';
import {AppComponent} from '../app.component';
import {HttpClient} from '@angular/common/http';
import {getAllBirdsResponse, getFoundBirds,Bird, PostData} from 'src/assets/components/components';
import {environment} from 'src/environments/environment';
import {FormControl} from '@angular/forms';
import {WikirestService} from '../services/wiki.service';
import {fromEvent,debounceTime, Subscription} from 'rxjs';
import {filter} from 'rxjs/operators';
import {MatInput} from '@angular/material/input';

@Component({
  selector: 'app-library',
  templateUrl: './library.component.html',
  styleUrls: ['./library.component.css'],
})

export class LibraryComponent implements OnInit{

  jsonUrl = 'assets/data.json';
  alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split('');
  searchInput: FormControl = new FormControl();
  pageSearch: FormControl = new FormControl();
  selectedOption: FormControl = new FormControl('');
  currentPageNumber: Number = 0;
  showNavButtonsFoundFilter = true;
  showNavButtonsLetterFilter = true;
  disableShowFoundFilter = false;
  showNothingFoundError: Boolean = false;
  isLoading: boolean = false;
  //nrOfPages is set dynamically but there is a bug when returning from species if initially set too small
  nrOfPages = 100000;
  //lengthOfSet is hardcoded to be static 30
  lenghtOfSet = 30;
  lengthOfBirds = 0;
  private pageSearchSubscription: Subscription | undefined;
  @ViewChild(MatInput) matInput!: MatInput;
  
  constructor(
    private router: Router,
    public mainApp: AppComponent,
    private http: HttpClient,
    private wikiRest: WikirestService,
    private route: ActivatedRoute
    ) {
  }

  ngOnInit() {
    this.getAllBirds();
    this.getSetOfBirds(0);
    this.getYourFoundBirds();
    this.setOfBirdsBackup.data = this.setOfBirds.data;

    this.route.queryParams.subscribe(params => {
      if(params['imagePage']){
        const numericValue = parseInt(params['imagePage'], 10);
        this.currentPageNumber = numericValue.valueOf();
      }
    });

    this.selectedOption.valueChanges.subscribe(value => {
      this.filterByLetter(value);
    });

    this.searchInput.valueChanges
    .pipe(
      debounceTime(300)
      )
      .subscribe(() => {
        const value = this.searchInput.value;
        if (value === "") {
          this.showNavButtonsFoundFilter = true;
          this.disableShowFoundFilter = false;
        } else {
          this.showNavButtonsFoundFilter = false;
          this.disableShowFoundFilter = true;
        }
        this.getSearchSet(value);
      });

      this.pageSearch.valueChanges.subscribe(value => {
      if(value == ''){
      }else{
        const numericValue = parseInt(value, 10);
        this.currentPageNumber = numericValue - 1;
        if (Number.isNaN(this.currentPageNumber.valueOf())){
          this.currentPageNumber = 0;
        }else if(this.currentPageNumber.valueOf()>=this.nrOfPages.valueOf()){
          this.currentPageNumber = this.nrOfPages.valueOf()-1;
        }
        
        this.changePage(0);}
    });
    this.changePage(0);
  }


  navigateToSpecies(imageId: string, imageName: string,imageSound:string, imageDesc: string, imageGenus:boolean): void {
    this.router.navigate(['species-page'], {
      queryParams: {
        imageId: encodeURIComponent(imageId),
        imageName: encodeURIComponent(imageName),
        imageSound: encodeURIComponent(imageSound),
        imageDesc: encodeURI(imageDesc),
        imageGenus: imageGenus,
        imagePage: this.currentPageNumber
      }
    });
  }

  setOfBirds: getAllBirdsResponse = {
    data:[],
    timestamp: ""
  }

  allBirds: getAllBirdsResponse = {
    data:[],
    timestamp: ""
  }

  setOfBirdsBackup: getAllBirdsResponse = {
    data:[],
    timestamp: ""
  }

  yourFoundBirds: getFoundBirds = {
    data:[],
    timestamp: ""
  }

  //filters displayed birds by the chosen first letter
  filterByLetter(selectedValue: any) {
    if (selectedValue == ""){
      this.showNavButtonsLetterFilter = true;
      this.getSetOfBirds(this.currentPageNumber);
    }else{
      this.showNavButtonsLetterFilter = false;
      this.setOfBirds.data = this.allBirds.data.filter(card => card.Name.startsWith(selectedValue));
      for (let i = 0; i < this.setOfBirds.data.length; i++) {
        this.setDataImageToWikiImage(this.getWikiLinkTitle(i),i);
      }
    }
  }

  //filters displayed birds by current users found birds list
  filterByFound(shouldFilter: boolean): void {
    this.showNavButtonsFoundFilter = !shouldFilter;
    this.setOfBirds.data = this.allBirds.data;

    if (shouldFilter) {
      if (this.yourFoundBirds.data.length == 0) {
        this.showNothingFoundError = true;
      } 
      else {
        const templist: Bird[] = [];
        const uniqueBirdIdsSet = new Set<string>();
        for (let i = 0; i < this.yourFoundBirds.data.length; i++){
          if (!uniqueBirdIdsSet.has(this.yourFoundBirds.data[i].birdId)){
            uniqueBirdIdsSet.add(this.yourFoundBirds.data[i].birdId);
            const foundBird = this.allBirds.data.find(bird => bird.Id === this.yourFoundBirds.data[i].birdId);
            if(foundBird){
              templist.push(foundBird);
              
            }
          }
        }
        this.setOfBirds.data = templist;
        for(let i = 0; i<this.setOfBirds.data.length;i++){
          this.setDataImageToWikiImage(this.getWikiLinkTitle(i),i);
        }
      }
    } 
    else {
      this.showNothingFoundError = false;
      this.getSetOfBirds(this.currentPageNumber);
    }
  }

  //The isLoading variable will toggle a spinner when loading images between clicks on the arrows,
  //for proper functionality it ought to be covering the loading of images instead of preventing them in html
  async getSetOfBirds(pageNumber:Number){
    //this.isLoading = true;
    this.sendGetSetOfBirdsRequest(pageNumber).subscribe(
      (response: getAllBirdsResponse) => {
        this.setOfBirds = response;
        this.lenghtOfSet = response.data.length;
        for (let i = 0; i <= this.setOfBirds.data.length; i++) {
          this.setDataImageToWikiImage(this.getWikiLinkTitle(i),i);
        }
        this.isLoading = false;
      },
      err => {
        this.isLoading = false;
        console.error("Failed at sending data:" + err); 
      }
    );
    setTimeout(()=>{
      this.isLoading = false;
    },300)
  }

  sendGetSetOfBirdsRequest(pageNumber:Number) {
    return this.http.get<getAllBirdsResponse>(environment.identifyRequestURL+"/birds/list?set="+pageNumber);
  }

  //gets all birds stored in backend
  getAllBirds() {
    this.sendGetSetOfBirdsRequest(-1).subscribe(
      (response: getAllBirdsResponse) => {
        this.allBirds = response;
        this.lengthOfBirds=response.data.length;
        this.nrOfPages = Math.ceil(this.lengthOfBirds.valueOf()/this.lenghtOfSet.valueOf());
      },
      err => { 
        console.error("Failed at sending data:" + err); 
      }
    );
  }

  //gets the users found birds list
  getYourFoundBirds() {
    this.sendGetYourFoundBirdsRequest().subscribe(
      (response: getFoundBirds) => {
        this.yourFoundBirds = response
      },
      err => {
        console.error("Failed at getting all found birds" + err);
      }
    )
  }

  sendGetYourFoundBirdsRequest() {
    const userId = localStorage.getItem("userId");
    return this.http.get<getFoundBirds>(environment.identifyRequestURL+"/users/"+userId+"/birds/list");
  }

  //sends request to backend to search through birdlist
  getSearchSet(searchQuery : string) {
    this.sendGetSearchSetRequest(searchQuery).subscribe(
      (response: getAllBirdsResponse) => {
        this.setOfBirds = response
        for (let i = 0; i <= this.setOfBirds.data.length; i++) {
          this.setDataImageToWikiImage(this.getWikiLinkTitle(i),i);
        }
      },
      err => {
        console.error("Failed at getting all found birds" + err);
      }
    )
  }

  sendGetSearchSetRequest(searchQuery : string) {
    return this.http.get<getAllBirdsResponse>(environment.identifyRequestURL+"/birds/list?search="+searchQuery);
  }

  //increments displayed page by +/- one
  changePage(increment:Number) {
    if (increment.valueOf() < 0) {
      this.pageSearch.setValue('');
      this.currentPageNumber = this.currentPageNumber.valueOf() - 1;
    } else if (increment.valueOf() > 0) {
      this.pageSearch.setValue('');
      this.currentPageNumber = this.currentPageNumber.valueOf() + 1;
    }

    if (this.currentPageNumber.valueOf() < 0) {
      this.currentPageNumber = 0;
    } else if(this.currentPageNumber.valueOf()>=this.nrOfPages.valueOf()){
      this.currentPageNumber = this.nrOfPages.valueOf()-1;
    } else {
      this.getSetOfBirds(this.currentPageNumber);
    }
  }

  getWikiLinkTitle(index:number) {
    let cutOffIndex = this.setOfBirds.data[index].Description.indexOf('wiki/');
    let cutString = this.setOfBirds.data[index].Description.substring(cutOffIndex + 'wiki/'.length)
    return cutString;
  }

  async setDataImageToWikiImage(wikiTitle:string, index:number) {
    this.wikiRest.getWiki(wikiTitle).subscribe(data => {
      if(data.extract) {
        if(data.originalimage?.source && !data.originalimage.source.includes('map')) {
          this.setOfBirds.data[index].Image = data.originalimage?.source;
        }else{
          this.setOfBirds.data[index].Image = "assets/no_img_available.png"
        }
      }
    },
    err => { 
      console.error('something went wrong' + err) 
    }); 
  }
  
  getCurrentPage(){
    return this.currentPageNumber.valueOf() + 1;
  }
}
