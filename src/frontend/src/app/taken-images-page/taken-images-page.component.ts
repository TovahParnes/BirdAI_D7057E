import {Component} from '@angular/core';
import {SocialAuthService} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import {AppComponent} from '../app.component';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {environment} from 'src/environments/environment';
import {DeleteResponse, LoginUser,UpdateResponse,UserResponse, listOutput} from 'src/assets/components/components';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';

@Component({
  selector: 'app-taken-images-page',
  templateUrl: './taken-images-page.component.html',
  styleUrls: ['./taken-images-page.component.css']
})

export class TakenImagesPageComponent {

  private jsonUrl = 'assets/data.json';
  dataList: any[] = [];
  activeSubMenuIndex: number | null = null;
  openForm: boolean = false;
  updateDetailsForm!: FormGroup;
  userMe!: LoginUser;
  userList: listOutput = {
    data: [],
    timestamp: ""
  }

  constructor(
    private router: Router, 
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private formBuilder: FormBuilder,
    private http: HttpClient
    ) {
  }

  ngOnInit(): void {
    this.getData().subscribe((response) => {
      const data = response;
      this.dataList = data.find((item) => 'dataList' in item)?.dataList || [];
    });

    this.updateDetailsForm = this.formBuilder.group({
      postId: ['', Validators.required],
      birdId: ['', Validators.required],
      location: ['', Validators.required],
      comment: ['', Validators.required],
    });

    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.getCurrentUser(authKey).subscribe(
        (response: UserResponse) => {
          this.userMe = response.data;
          //after getting currentuser I have to immediatly run the getCurrentUserList or else the nginit will run this part before for some reason, 
          //the value of this.userMe is set properly outside nginit but not inside if it is not nestled like this
          this.getCurrentUserList().subscribe(
            (response: listOutput) => {
              this.userList.data = response.data;
              console.log(this.userList.data);
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

  parseDate(date:string){
    const temp = date.split("T");
    const newDate = temp[0];
    return(newDate);
  }

  toggleSubMenu(index: number, event: Event) {
    event.stopPropagation();
    this.activeSubMenuIndex = this.activeSubMenuIndex === index ? null : index;
  }

  updateForm(postId: string, location: string, birdId: string){
    document.body.style.overflow = 'hidden';
    this.updateDetailsForm = this.formBuilder.group({
      postId: [{ value: postId}, Validators.required],
      birdId: [{ value: birdId}, Validators.required],
      location: [{ value: location, disabled: false } , Validators.required],
      comment: [ '' , Validators.required],
    });
    this.openForm = true;
  }

  closePostForm(){
    document.body.style.overflow = 'auto';
    this.openForm = false;
    this.updateDetailsForm.reset();
  }

  getPostIdToDelete(id: string){
    this.deletePost(id);
    window.location.reload();
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

  getCurrentUser(token: string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    return this.http.get<UserResponse>(environment.identifyRequestURL+"/users/me",{ headers: header });
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

  updatePost(){
    const authKey = localStorage.getItem("auth");
    if(authKey){
      let postId = this.updateDetailsForm.get('postId')?.value?.value;
      let birdId = this.updateDetailsForm.get('birdId')?.value?.value;
      let locat = this.updateDetailsForm.get('location')?.value ?? "unknown";
      
      this.sendUpdate(authKey, postId, locat, birdId).subscribe(
        (response: UpdateResponse) => {
          this.openForm = false;
          document.body.style.overflow = 'auto';
          window.location.reload();
        },err => { 
          console.error("Failed at updating post with id: "+ postId + " " + err); 
        }
      )
    }
  }

  sendUpdate(token: string, postId: string, newPostValues: string, birdId: string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    const location = newPostValues;
    const body = {
      "birdId": birdId,
      location
    }
    return this.http.patch<UpdateResponse>(environment.identifyRequestURL+"/posts/"+postId, body, { headers: header });
  }
}
