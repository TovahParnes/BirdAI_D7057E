import {Component, ViewChildren, QueryList} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider, SocialUser} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';
import { Card2Component } from '../card/card.component';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import {LoginUser,UserResponse, listOutput} from 'src/assets/components/components';

@Component({
  selector: 'app-taken-images-page',
  templateUrl: './taken-images-page.component.html',
  styleUrls: ['./taken-images-page.component.css']
})

export class TakenImagesPageComponent {

  list1: any[] = [];
  list2: any[] = [];
  private jsonUrl = 'assets/data.json';

  // cardlist = [
  //   {title: 'Duck',imageSrc: 'assets/duck.jpg', date:'2023-10-05'},
  //   {title: 'Budgie',imageSrc: 'assets/undulat.jpg', date:'2023-10-04'},
  // ]

  constructor(
    private router: Router, 
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private http: HttpClient,
    ) {
  }
  userMe!: LoginUser;
  userList!: listOutput;
  


  navigateToSpecies(imageId: string, imageName: string, imageDate: string): void {
    this.router.navigate(['species-page'], {
      queryParams: {
        imageId: encodeURIComponent(imageId),
        imageName: encodeURIComponent(imageName),
        imageDate: encodeURIComponent(imageDate),
      }
      });
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
      this.list1 = data.find((item) => 'list1' in item)?.list1 || [];
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
