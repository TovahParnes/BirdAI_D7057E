import {Component, ViewChildren, QueryList} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider, SocialUser} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';
import { CardComponent } from '../card/card.component';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import {LoginUser,UserResponse, listOutput} from 'src/assets/components/components';

@Component({
  selector: 'app-taken-images-page',
  templateUrl: './taken-images-page.component.html',
  styleUrls: ['./taken-images-page.component.css'],
})

export class TakenImagesPageComponent {

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
    ) {
  }
  userMe!: LoginUser;
  userList!: listOutput;

  

  logout(): void {
    this.socialAuthService.signOut().then(() => this.router.navigate(['login']));
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

  toggleTheme(): void {
    //this.mainApp.switchDarkmodeSetting();
  }

  navigateToSpecies(imageId: string, imageName: string): void {
    this.router.navigate(['species-page'], {
      queryParams: {
        imageId: encodeURIComponent(imageId),
        imageName: encodeURIComponent(imageName),
      }
      });
  }

  //sorts by name but should prob not be visible as a button but rather done autmatically in the background
  sortCards() {
    this.cardlist.sort((a, b) => a.title.localeCompare(b.title));
  }

  filterCards(letter: string): void {
    if(this.found == true){
      this.cardlist = this.foundlist
    }else{
      this.cardlist = this.backupCards
    }
    this.cardlist = this.cardlist.filter(card => card.title.startsWith(letter));
    this.lastLetter = letter
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

  // getBackendData() {
  //   return this.http.get<any[]>(environment.identifyRequestURL+"/birds/list", this.userMe._id)
  // }

  getCurrentUser(token: string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    return this.http.get<UserResponse>(environment.identifyRequestURL+"/users/me",{ headers: header });
  }

  ngOnInit(): void {
    this.getData().subscribe((response) => {
      const data = response;
      this.cardlist = data.find((item) => 'list2' in item)?.list2 || [];
      this.foundlist = data.find((item) => 'list1' in item)?.list1 || [];
      this.backupCards = this.cardlist
    });
    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.getCurrentUser(authKey).subscribe(
        (response: UserResponse) => {
          console.log("Succesfully got userMe");
          console.log(response.data);
          this.userMe = response.data;
          console.log(this.userMe);
          //after getting currentuser I have to immediatly run the getCurrentUserList or else the nginit will run this part before for some reason, 
          //the value of this.userMe is set properly outside nginit but not inside if it is not nestled like this
          this.getCurrentUserList().subscribe(
            (response: listOutput) => {
              console.log("Succesfully retrieved user list");
              console.log(response.data);
              this.userList = response;
            },err => { 
              console.error("Failed at getting user list:" + err); 
            }
          )
          console.log(this.userList)
        },err => { 
          console.error("Failed at getting userMe:" + err); 
        }
      )
    }
  }
  getCurrentUserList(){
    return this.http.get<listOutput>(environment.identifyRequestURL+"/users/"+this.userMe._id+"/posts/list");
  }
}
