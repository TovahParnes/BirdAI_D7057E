import {Component, ViewChildren, QueryList, HostListener, AfterViewInit} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider, SocialUser} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';
import { Card2Component} from '../card/card.component';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import {DeleteResponse, LoginUser,UpdateResponse,UserResponse, listOutput} from 'src/assets/components/components';
import { Location } from '@angular/common'
import { Block } from '@angular/compiler';

@Component({
  selector: 'app-taken-images-page',
  templateUrl: './taken-images-page.component.html',
  styleUrls: ['./taken-images-page.component.css']
})

export class TakenImagesPageComponent {

  list1: any[] = [];
  list2: any[] = [];
  private jsonUrl = 'assets/data.json';
  dropDownVisibility = false;

  // cardlist = [
  //   {title: 'Duck',imageSrc: 'assets/duck.jpg', date:'2023-10-05'},
  //   {title: 'Budgie',imageSrc: 'assets/undulat.jpg', date:'2023-10-04'},
  // ]

  constructor(
    private router: Router, 
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private http: HttpClient,
    private location: Location,
    ) {
  }
  userMe!: LoginUser;
  userList: listOutput = {
    data: [],
    timestamp: ""
  }

  parseDate(date:string){
    const temp = date.split("T");
    const newDate = temp[0];
    return(newDate);
  }

  getDataToUpdate(postId:string,birdId:string){
    const textField = document.getElementById("updateTextField") as HTMLInputElement;
    this.updatePost(postId,textField.value,birdId);
    this.toggleUpdatePostField("null");
  }

  getPostIdToDelete(){
    const textField = document.getElementById("postIdTextField") as HTMLInputElement;
    this.deletePost(textField.value)
  }
  
  getCurrentUserListData(){
    console.log(this.userList.data)
  }

  navigateToSpecies(imageId: string, imageName: string, imageDate: string, imageDesc: string): void {
    this.router.navigate(['species-page'], {
      queryParams: {
        imageId: encodeURIComponent(imageId),
        imageName: encodeURIComponent(imageName),
        imageDate: encodeURIComponent(imageDate),
        imageDesc: encodeURIComponent(imageDesc),
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
          this.userMe = response.data;
          localStorage.setItem("userId",response.data._id);
          //after getting currentuser I have to immediatly run the getCurrentUserList or else the nginit will run this part before for some reason, 
          //the value of this.userMe is set properly outside nginit but not inside if it is not nestled like this
          this.getCurrentUserList().subscribe(
            (response: listOutput) => {
              this.userList.data = response.data;
              console.log(this.userList.data)
            },err => { 
              console.error("Failed at getting user list:" + err); 
            }
          )
        },err => { 
          console.error("Failed at getting userMe:" + err); 
        }
      )
    }
  }
  getCurrentUserList(){
    return this.http.get<listOutput>(environment.identifyRequestURL+"/users/"+this.userMe._id+"/posts/list");
  }

  deletePost(postId: string){
    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.sendDelete(authKey,postId).subscribe(
        (response: DeleteResponse) => {
          window.location.reload();
        },err => { 
          console.error("Failed at deleting post with id: "+ postId + " " + err); 
        }
      )
    }
  }

  sendDelete(token: string, postId:string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    return this.http.delete<DeleteResponse>(environment.identifyRequestURL+"/posts/"+postId,{ headers: header });
  }

  updatePost(postId:string, newPostValues:string,birdId:string){
    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.sendUpdate(authKey,postId, newPostValues, birdId).subscribe(
        (response: UpdateResponse) => {
        },err => { 
          console.error("Failed at updating post with id: "+ postId + " " + err); 
        }
      )
    }
  }

  sendUpdate(token: string, postId: string,newPostValues:string,birdId:string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    const location = newPostValues;
    const body = {
      "birdId": birdId,
      location
    }
    return this.http.patch<UpdateResponse>(environment.identifyRequestURL+"/posts/"+postId,body,{ headers: header });
  }

  // @HostListener('document:click', ['$event'])
  // handleDocumentClick(event: MouseEvent) {
  //   console.log("clicked");
  //   const dropdownElement = document.getElementById('dropdown-menu');
  //   const elements = document.querySelectorAll('dropdown');
  //   console.log(elements);
  //   // if (this.dropDownVisibility) {
  //   //   const dropdownElement = document.getElementById('dropdown-menu');
  //   //   if (dropdownElement && !dropdownElement.contains(event.target as Node)) {
  //   //     this.dropDownVisibility = false;
  //   //   }
  //   // }
  // }

  toggleUpdatePostField(id:string){
    const updatePostField = document.getElementById(id);
    if(updatePostField){
      if(updatePostField.style.display == 'none'){
        updatePostField.style.display='block';
      }else{
        updatePostField.style.display='none';
      }
    }
  }
  
}


