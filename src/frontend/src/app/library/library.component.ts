import {Component, OnInit} from '@angular/core';
import {Router, ActivatedRoute} from '@angular/router';
import {AppComponent} from '../app.component';
import {HttpClient} from '@angular/common/http';
import {getAllBirdsResponse, getFoundBirds} from 'src/assets/components/components';
import {environment} from 'src/environments/environment';
import {FormControl} from '@angular/forms';
import {WikirestService} from '../services/wiki.service';

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
  currentPageNumber: Number = 0;
  showNavButtons = true;
  disableShowFoundFilter = false;
  showNothingFoundError: Boolean = false;
  isLoading: boolean = false;
  nrOfPages = 1;
  //lengthOfSet is hardcoded to be static 30
  lenghtOfSet = 30;
  lengthOfBirds = 0;
  
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

    this.route.queryParams.subscribe(params => {
      if(params['imagePage']){
        const numericValue = parseInt(params['imagePage'], 10);
        this.currentPageNumber = numericValue;
      }
    });

    this.selectedOption.valueChanges.subscribe(value => {
      this.filterByLetter(value);
    });

    this.searchInput.valueChanges.subscribe(value => {
      if (value == "") {
        this.showNavButtons = true;
        this.disableShowFoundFilter = false;
      } else {
        this.showNavButtons = false;
        this.disableShowFoundFilter = true;
      }
      this.getSearchSet(value);
    });

    this.pageSearch.valueChanges.subscribe(value => {
      const numericValue = parseInt(value, 10);
      this.currentPageNumber = numericValue - 1;
      if (Number.isNaN(this.currentPageNumber.valueOf())){
        this.currentPageNumber = 0;
      }else if(this.currentPageNumber.valueOf()>=this.nrOfPages.valueOf()){
        this.currentPageNumber = this.nrOfPages.valueOf()-1;
      }
      this.changePage(0);
    });
    this.changePage(0);
  }

  navigateToSpecies(imageId: string, imageName: string,imageSound:string, imageDesc: string, imageGenus:Boolean): void {
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

  filterByLetter(selectedValue: any) {
    this.setOfBirds.data = this.setOfBirdsBackup.data;
    this.setOfBirds.data = this.setOfBirds.data.filter(card => card.Name.startsWith(selectedValue));
  }

  filterByFound(shouldFilter: boolean): void {
    this.showNavButtons = !shouldFilter;
    this.setOfBirds.data = this.allBirds.data;

    if (shouldFilter) {
      if (this.yourFoundBirds.data.length == 0) {
        this.showNothingFoundError = true;
      } 
      
      else {
        for (let i = 0; i < this.yourFoundBirds.data.length; i++) {
          this.setOfBirds.data = this.setOfBirds.data.filter(item => item.Id.includes(this.yourFoundBirds.data[i].birdId));
          this.setDataImageToWikiImage(this.getWikiLinkTitle(i),i);
        }
      }
    } 
    
    else {
      this.showNothingFoundError = false;
      this.setOfBirds.data = this.setOfBirdsBackup.data;
    }
  }


  async getSetOfBirds(pageNumber:Number){
    this.isLoading = true;
    this.sendGetSetOfBirdsRequest(pageNumber).subscribe(
      (response: getAllBirdsResponse) => {
        this.setOfBirds = response;
        this.setOfBirdsBackup.data = response.data;
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

  changePage(increment:Number) {
    if (increment.valueOf() < 0) {
      this.currentPageNumber = this.currentPageNumber.valueOf() - 1;
    } else if (increment.valueOf() > 0) {
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
        if(data.originalimage?.source) {
          this.setOfBirds.data[index].Image = data.originalimage?.source;
          this.setOfBirdsBackup.data[index].Image = data.originalimage?.source;
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
