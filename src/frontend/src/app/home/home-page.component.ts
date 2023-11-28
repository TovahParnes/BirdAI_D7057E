import {Component, OnInit} from '@angular/core';
import {SocialAuthService} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import {AnalyzeResponse, PostData, AnalyzedBird, UserBirdList, AdminResponse, UserResponse} from 'src/assets/components/components';
import {AppComponent} from '../app.component';
import {FormBuilder, FormGroup, FormControl, Validators} from '@angular/forms';
import {HttpClient} from '@angular/common/http';
import {environment} from 'src/environments/environment';
import {Observable, catchError} from 'rxjs';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css'],
})


export class MainPageComponent implements OnInit {

  form!: FormGroup;
  postDetailsForm!: FormGroup;
  selectedImage: any;
  isLoading: boolean = false;
  createPostForm: boolean = false;
  latestAnalyzedBird!: AnalyzedBird
  analyzedBirdList: UserBirdList = {
    birds: []
  }
  data: any;
  dataImg: any;
  analyzed: AnalyzeResponse | null = null;
  togglePostView = false;
  toggleConfirmView = false;
  fileFormat = "";

  constructor(
    private router: Router,
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private formBuilder: FormBuilder,
    private http: HttpClient) {
  }

  ngOnInit() {
    this.getUserMe();
    this.getCurrentAdmin();
    this.form = this.formBuilder.group({
      option: new FormControl(),
    });
    this.postDetailsForm = this.formBuilder.group({
      name: ['', Validators.required],
      accuracy: ['', Validators.required],
      location: ['', Validators.required],
      comment: ['', Validators.required],
    });
    console.log(localStorage.getItem("auth"));
  }
  
  convertAccuracy(accuracy: string){
    return (Number(accuracy) * 100);
  }

  convertAccuracyToString(accuracy: string){
    return this.convertAccuracy(accuracy).toString()+"%";
  }

  onFileSelected(event: any) {
    const file = event.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = (e) => {
        this.selectedImage = reader.result;
        const dataUrl = this.selectedImage as string;
        console.log(dataUrl)
        const fileFormat = dataUrl.substring(dataUrl.indexOf('/') + 1, dataUrl.indexOf(';'));
        this.fileFormat = fileFormat
      };
    }
  }

  onClear() {
    this.selectedImage = null;
  }

  postImage(token: string): Observable<AnalyzeResponse> {
    const header = {
      'Authorization': `Bearer ${token}`
    };
    const body = {'data': `${this.selectedImage}`, 'fileType': this.fileFormat};
    return this.http.post<AnalyzeResponse>(environment.identifyRequestURL+"/ai/inputimage", body, { headers: header });
  }

  onSubmit(el: HTMLElement) {
    this.isLoading = true;
    const authKey = localStorage.getItem("auth");
    if(authKey){
    this.postImage(authKey).subscribe(
      (response: AnalyzeResponse) => {
        console.log("Succesfully sent data");
        console.log(response.data);
        this.dataImg = this.selectedImage;
        this.analyzed = response;
        if (this.analyzed.data.length == 0) {
          this.isLoading = false;
          console.log("No birds found");
        }
        else {
          this.isLoading = false;
          this.addNewBird(this.analyzed.data[0].aiBird.name, this.analyzed.data[0].userMedia.data, this.analyzed.data[0].aiBird.accuracy);
        }
      },
      err => {
        this.isLoading = false;
        console.error("Failed at sending data:" + err);
      }
    );
    this.togglePostView = true;
    el.scrollIntoView();
    }
  }

  addNewBird(name: string, imageUrl:string, accuracy:string){
    const newitem = {"title": name, "image": imageUrl, "accuracy": accuracy}
    this.analyzedBirdList.birds.push(newitem);
  }

  getLatestBird(){
    const len = this.analyzedBirdList.birds.length - 1;
    return(this.analyzedBirdList.birds[len]);
  }

  getCurrentDateAndTime() {
    const dateTime = new Date();
    return dateTime.toLocaleString();
  }

  resetForm(){
    this.form.reset();
    this.postDetailsForm.reset();
    this.analyzed = null;
    this.togglePostView = false;
    this.selectedImage = null;
    setTimeout(function() {
      window.scrollTo({ top: 0, behavior: 'smooth' });
    }, 100);
  }
  
  openPostForm(){
    if (this.analyzed){
      document.body.style.overflow = 'hidden';
      this.postDetailsForm = this.formBuilder.group({
        name: [{ value: this.analyzed.data[0].aiBird.name, disabled: true }, Validators.required],
        accuracy: [{ value: this.convertAccuracyToString(this.analyzed.data[0].aiBird.accuracy), disabled: true }, Validators.required],
        date: [{ value: this.getCurrentDateAndTime(), disabled: true }, Validators.required],
        location: [ '' , Validators.required],
        comment: [ '' , Validators.required],
      });
      this.createPostForm = true;
    }
  }

  closePostForm(){
    document.body.style.overflow = 'auto';
    this.createPostForm = false;
    this.postDetailsForm.reset();
  }

  sendPost(token:string) {
      if (this.analyzed){
        const header = {
          'Authorization': `Bearer ${token}`
        };

        let location = this.postDetailsForm.get('location')?.value;
        console.log(location);
        const postData = {'_id': "no", 'birdId': this.analyzed.data[0].birdId, 'location': location, "media":{'data': this.selectedImage, 'filetype': this.fileFormat}};
        console.log("Send:", postData);
        console.log(environment.identifyRequestURL+"/posts",postData,{ headers: header });
        return this.http.post<PostData>(environment.identifyRequestURL+"/posts",postData,{ headers: header });
      } else {
        return null
      }
  }

  createPost(){
    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.sendPost(authKey)?.subscribe(
        (response: PostData) => {
          console.log("Succesfully sent data");
          this.dataImg = this.selectedImage;
        },
        err => {
          console.error("Failed at sending data:" + err);
        }
        );
      setTimeout(() => {
        document.body.style.overflow = 'auto';
        this.router.navigate(['takenImages']);
      }, 100);
    }
    document.body.style.overflow = 'auto';
  }

  getCurrentAdmin(){
    try{
    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.sendGetCurrentAdmin(authKey).pipe(
        catchError((error: any) => {
          return [];
        })
      ).subscribe(
        (response: AdminResponse) => {
          localStorage.setItem("currentAdmin",response.data.user._id);
        }
      )
    }
  }catch(error){
  }
  }

  sendGetCurrentAdmin(token:string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    return this.http.get<AdminResponse>(environment.identifyRequestURL+"/admins/me",{ headers: header });
  }

  getUserMe(){
    const authKey = localStorage.getItem("auth");
    if(authKey){
      this.getCurrentUser(authKey).subscribe(
        (response: UserResponse) => {
          localStorage.setItem("userId",response.data._id);
    }
  )}
}
getCurrentUser(token: string){
  const header = {
    'Authorization': `Bearer ${token}`
  };
  return this.http.get<UserResponse>(environment.identifyRequestURL+"/users/me",{ headers: header });
}
  
}
