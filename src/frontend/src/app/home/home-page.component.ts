import {Component, OnInit, ViewChild} from '@angular/core';
import {SocialAuthService, SocialUser} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import {AppComponent} from '../app.component';
import { FormBuilder, FormGroup, FormControl } from '@angular/forms';
import { MatStepper } from '@angular/material/stepper';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';

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

  data: any;
  dataImg: any;

  constructor(
    private router: Router, 
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private formBuilder: FormBuilder,
    private httpClient: HttpClient) {
  }

  ngOnInit() {
      this.socialAuthService.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null);
    }),

    this.form = this.formBuilder.group({
      option: new FormControl(), // Initialize with a default value
    });
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

  onSubmit(el: HTMLElement) {
    this.isLoading = true;

    const header = {'Authorization': `Bearer ${environment.secret}`};
    const body = {'img': `${this.selectedImage}`};
    this.httpClient.post<any>(environment.identifyRequestURL, body, { headers: header })
    .subscribe(
      () => {
        console.log("Succesfully sent data");
        this.form.reset();
        this.dataImg = this.selectedImage;
        this.selectedImage = null;
        el.scrollIntoView();

        
      },
      err => { 
        console.error("Failed at sending data:" + err); 
      }
    );
    this.isLoading = false
  }
}




