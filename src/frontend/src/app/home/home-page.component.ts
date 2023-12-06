import {Component, OnInit} from '@angular/core';
import {SocialAuthService} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import {AnalyzeResponse, PostData, AnalyzedBird, UserBirdList, AdminResponse, UserResponse} from 'src/assets/components/components';
import {AppComponent} from '../app.component';
import {FormBuilder, FormGroup, FormControl, Validators} from '@angular/forms';
import {HttpClient} from '@angular/common/http';
import {environment} from 'src/environments/environment';
import {Observable, catchError} from 'rxjs';
import { WikiPageSegment, WikiSummary, WikirestService } from '../services/wiki.service';
import {Ng2ImgMaxService} from 'ng2-img-max'

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
  compressed_img: string = "";

  constructor(
    private router: Router,
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private formBuilder: FormBuilder,
    private http: HttpClient,
    private wikiRest: WikirestService,
    private ng2ImgMax: Ng2ImgMaxService) {
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
  }

  round(value: Number, precision: Number) {
    var multiplier = Math.pow(10, precision.valueOf() || 0);
    return Math.round(value.valueOf() * multiplier) / multiplier;
}
  
  convertAccuracy(accuracy: Number): Number {
    const newAccuracy = (accuracy.valueOf() * 100);
    return this.round(newAccuracy, 1);
  }

  convertAccuracyToString(accuracy: Number): string {
    return this.convertAccuracy(accuracy).toString()+"%";
  }

  onFileSelected(event: any) {
    const file = event.target.files[0];
    if (file) {
      const analyzeError = document.getElementById('analyzeError');
      if(analyzeError){
      analyzeError.style.display = 'none';
      }
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = (e) => {
        this.selectedImage = reader.result;
        if(this.selectedImage.length>15*1024*1024){
          if (analyzeError){
            analyzeError.style.display = 'block';
            this.selectedImage = false;
          }
        }else{
          const dataUrl = this.selectedImage as string;
          const fileFormat = dataUrl.substring(dataUrl.indexOf('/') + 1, dataUrl.indexOf(';'));
          this.fileFormat = fileFormat;
  
          const percentageReduction = 0.7;
          const targetFileSize = file.size * (1 - percentageReduction);
          const maxSizeInMB = targetFileSize * 0.000001;
          this.compressImage(file,maxSizeInMB);
        }
      };
    }
  }

  onClear() {
    this.selectedImage = null;
  }

  postImageForAnalysing(token: string): Observable<AnalyzeResponse> {
    const header = {
      'Authorization': `Bearer ${token}`
    };
    const body = {
      'data': `${this.selectedImage}`,
      'fileType': this.fileFormat
    };
    return this.http.post<AnalyzeResponse>(environment.identifyRequestURL+"/ai/inputimage", body, { headers: header });
  }

  onSubmit(el: HTMLElement) {
    this.isLoading = true;
    const authKey = localStorage.getItem("auth");
    if(authKey){
    this.postImageForAnalysing(authKey).subscribe(
      (response: AnalyzeResponse) => {
        this.dataImg = this.selectedImage; //TABORT???????
        this.analyzed = response;

        // No birds found
        if (this.analyzed.data.length == 0) {
          this.isLoading = false;
        }

        // Bird found
        else {
          this.isLoading = false;
          this.addNewBird(
            this.analyzed.data[0].aiBird.name,
            this.analyzed.data[0].description,
            this.analyzed.data[0].aiBird.accuracy
              );
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

  addNewBird(name: string, imageUrl: string, accuracy: Number){
    const newitem = {"title": name, "image": imageUrl, "accuracy": accuracy}
    this.analyzedBirdList.birds.push(newitem);
    const len = this.analyzedBirdList.birds.length - 1;
    const wikiLink = this.getWikiLinkTitle(len);
    this.setDataImageToWikiImage(wikiLink,len);
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
        let comment = this.postDetailsForm.get('comment')?.value;

        const postData = {
          'accuracy': this.analyzed.data[0].aiBird.accuracy,
          'birdId': this.analyzed.data[0].birdId,
          'comment': comment,
          'location': location,
          "media":{
            'data': this.compressed_img,
            'filetype': this.fileFormat
          }
        };

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

getWikiLinkTitle(index:number){
  let cutOffIndex = this.analyzedBirdList.birds[index].image.indexOf('wiki/');
  let cutString = this.analyzedBirdList.birds[index].image.substring(cutOffIndex + 'wiki/'.length)
  return cutString;
}

async setDataImageToWikiImage(wikiTitle:string,index:number){
  this.wikiRest.getWiki(wikiTitle).subscribe(data => {
    if(data.extract){
    //this.wikiData = data;
    if(data.originalimage?.source){
      this.analyzedBirdList.birds[index].image = data.originalimage?.source;
      this.analyzedBirdList.birds[index].image = data.originalimage?.source;
    }
    }
  }, err => { console.log('something went wrong' + err)
}); 

}

compressImage(file: File, maxSizeInMB: number)  {
  this.ng2ImgMax.compressImage(file, maxSizeInMB)
    .subscribe(compressedImage => {
      this.blobToBase64(compressedImage).then((result:string)=>{
        this.compressed_img = result;
      }
      );
    }, error => {
      console.log(error.reason);
   });
}

blobToBase64(blob:Blob) {
  return new Promise<string>((resolve, _) => {
    const reader = new FileReader();
    reader.onloadend = () => resolve(reader.result as string);
    reader.readAsDataURL(blob);
  });
}

}
