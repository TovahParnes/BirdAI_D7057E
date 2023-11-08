import {Component, OnInit } from '@angular/core';
import {SocialAuthService } from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import {AnalyzeResponse, PostData, AnalyzedBird, UserBirdList} from 'src/assets/components/components';
import {AppComponent} from '../app.component';
import { FormBuilder, FormGroup, FormControl } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css'],
})

export class MainPageComponent implements OnInit {

  isLinear = false;
  form!: FormGroup;
  selectedImage: any;
  isLoading: boolean = false;
  latestAnalyzedBird!: AnalyzedBird
  analyzedBirdList: UserBirdList = {
    birds: []
  }
  data: any;
  dataImg: any;
  analyzed: AnalyzeResponse | null = null;
  togglePostView = true;
  toggleConfirmView = false;
  fileFormat = "";

  constructor(
    private router: Router,
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private formBuilder: FormBuilder,
    private http: HttpClient) {
  }

  convertAccuracy(accuracy: string){
    return (Number(accuracy) * 100).toString()+"%";
  }

  ngOnInit() {
    this.form = this.formBuilder.group({
      option: new FormControl(),
    });
    console.log(localStorage.getItem("auth"));
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
        this.form.reset();
        this.dataImg = this.selectedImage;
        this.selectedImage = null;
        el.scrollIntoView();
        this.analyzed = response;
        if (this.analyzed.data.length == 0) {
          console.log("No birds found");
        }
        else {
          this.addNewBird(this.analyzed.data[0].aiBird.name, this.analyzed.data[0].userMedia.data, this.analyzed.data[0].aiBird.accuracy);
        }
      },
      err => {
        console.error("Failed at sending data:" + err);
      }
    );
    this.isLoading = false;
    this.togglePostView = true;
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

  togglePreview(){
    if(this.togglePostView){
      this.togglePostView = false;
    }else{
      this.togglePostView = true;
    }
  }

  sendPost(token:string) {
      if (this.analyzed){
      const header = {
        'Authorization': `Bearer ${token}`
      };
      const postData = {'_id': "no", 'birdId': this.analyzed.data[0].birdId,'location': "no", "media":{'data':this.dataImg, 'filetype': this.fileFormat }};
      console.log(postData);
      return this.http.post<PostData>(environment.identifyRequestURL+"/posts",postData,{ headers: header });
      }else{
        return null
      }
  }

  createPost(){
    const authKey = localStorage.getItem("auth");
    if(authKey){
    this.sendPost(authKey)?.subscribe(
      (response: PostData) => {
        console.log("Succesfully sent data");
        console.log(response);
        this.form.reset();
        this.dataImg = this.selectedImage;
        this.selectedImage = null;
        this.togglePreview();
        this.toggleConfirmView = true;
      },
      err => {
        console.error("Failed at sending data:" + err);
      }
    );
    }
  }
}
