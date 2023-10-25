import {Component, OnInit, ViewChild} from '@angular/core';
import {SocialAuthService, SocialUser} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import {AnalyzeResponse} from 'src/assets/components/components';
import {AppComponent} from '../app.component';
import { FormBuilder, FormGroup, FormControl } from '@angular/forms';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Observable } from 'rxjs';
import { NgOptimizedImage } from '@angular/common'
import jsonData from '../../assets/data.json';

interface list3 {
  list3:{   
    //dataitem: dataitem  
    title: string;
    image: string;
    accuracy: string;
  }[]
}

interface dataitem {
  title: string;
  image: string;
  accuracy: string;
}

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css'],
})

export class MainPageComponent implements OnInit {

  user: SocialUser = new SocialUser;
  loggedIn: boolean = false;
  isLinear = false;
  form!: FormGroup;
  selectedImage: any;
  isLoading: boolean = false;
  print =""
  element!: dataitem
  dummyitem: dataitem = {title:'test',image:'test',accuracy:'test'}
  jsonlist = jsonData[2] as list3;
  //jsonlist: list3 ={ list3: [{dataitem: this.dummyitem},]}

  data: any;
  dataImg: any;
  analyzed: AnalyzeResponse | null = null;


  constructor(
    private router: Router, 
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private formBuilder: FormBuilder,
    private http: HttpClient) {
  }

  ngOnInit() {
      this.socialAuthService.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null);
    }),

    this.form = this.formBuilder.group({
      option: new FormControl(), // Initialize with a default value
    });
    this.getJsonData();
  }

  onFileSelected(event: any) {
    const file = event.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = (e) => {
        this.selectedImage = reader.result;
      };
    }
  }

  onClear() {
    this.selectedImage = null;
  }

  postImage(): Observable<AnalyzeResponse> {
    //console.log(header);
    const body = {'data': `${this.selectedImage}`, 'fileType': "JPG"};
    return this.http.post<AnalyzeResponse>(environment.identifyRequestURL+"/ai/inputimage", body);
  }

  onSubmit(el: HTMLElement) {
    this.isLoading = true;

    this.postImage().subscribe(
      (response: AnalyzeResponse) => {
        console.log("Succesfully sent data");
        console.log(response.data);
        this.form.reset();
        this.dataImg = this.selectedImage;
        this.selectedImage = null;
        el.scrollIntoView();
        this.analyzed = response;
        this.addNewBird(this.analyzed.data.name, this.analyzed.data.picture.data, this.analyzed.data.accuracy);
      },
      err => { 
        console.error("Failed at sending data:" + err); 
      }
    );
    this.isLoading = false
  }

  //AddnewBird sparar temporärt i jsonlist, detta skulle kunna skickas mellan sidorna
  addNewBird(name: string, imageUrl:string, accuracy:string){
    //const newitem : { dataitem: dataitem } = {dataitem: {"title": name, "image": imageUrl, "accuracy": accuracy} };
    const newitem = {"title": name, "image": imageUrl, "accuracy": accuracy}
    this.jsonlist.list3.push(newitem)
    //this.jsonlist.list3.push(newitem)

  }

  //exempel på hur det ifrån json
  getJsonData(){
    // if (this.jsonlist) {
    //   this.element = this.jsonlist.list3[1].dataitem;
    // }
    this.element = this.jsonlist.list3[1]
  }


}