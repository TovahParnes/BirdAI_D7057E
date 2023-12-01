import {Component, OnInit} from '@angular/core';
import {SocialAuthService} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import {AppComponent} from '../app.component';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {getAllBirdsResponse} from 'src/assets/components/components';
import {environment} from 'src/environments/environment';
import {FormControl} from '@angular/forms';

@Component({
  selector: 'app-library',
  templateUrl: './library.component.html',
  styleUrls: ['./library.component.css'],
})

export class LibraryComponent implements OnInit {

  jsonUrl = 'assets/data.json';
  alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split('');
  searchInput: FormControl = new FormControl();
  selectedOption: FormControl = new FormControl('');
  cardlist: any[] = [];
  foundlist: any[] = [];
  
  constructor(
    private router: Router,
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private http: HttpClient,
    ) {
  }

  ngOnInit() {
    this.getAllBirds();
    this.cardlist = this.allBirds.data;
    this.foundlist = this.allBirds.data;

    this.selectedOption.valueChanges.subscribe(value => {
      this.filterByLetter(value);
    });

    this.searchInput.valueChanges.subscribe(value => {
      this.filterBySearch(value.toUpperCase());
    });
  }

  navigateToSpecies(imageId: string, imageName: string, imageDesc: string): void {
    this.router.navigate(['species-page'], {
      queryParams: {
        imageId: encodeURIComponent(imageId),
        imageName: encodeURIComponent(imageName),
        imageDesc: encodeURI(imageDesc)
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

  getAllBirds(){
    this.sendRequestGetBirds().subscribe(
      (response: getAllBirdsResponse) => {
        this.allBirds = response;
        this.allBirdsBackup.data = response.data;
      },
      err => { 
        console.error("Failed at sending data:" + err); 
      }
    );
  }

  sendRequestGetBirds() {
    return this.http.get<getAllBirdsResponse>(environment.identifyRequestURL+"/birds/list?set=0");
  }
}
