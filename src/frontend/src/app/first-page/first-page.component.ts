import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import {GoogleLoginProvider, SocialAuthService} from '@abacritt/angularx-social-login';
import { AppComponent } from '../app.component';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { UserResponse } from 'src/assets/components/components';

declare const gapi: any;

@Component({
  selector: 'app-first-page',
  templateUrl: './first-page.component.html',
  styleUrls: ['./first-page.component.css']

})

export class FirstPageComponent {

  constructor(
    public mainApp: AppComponent,
    private router: Router,
    private httpClient: HttpClient,
    private socialAuthService: SocialAuthService,
    ) {
  }

  public triedLogIn = false;
  imgsize = 150

  postLoggedInUser(): Observable<UserResponse> {
    const username = `${this.mainApp.user.name}`.replace(/[\s!@#$%^&*()_+{}\[\]:;<>,.?~\\|/`'"-]/g, '')
    const body = {
      'username': username,'authId': `${this.mainApp.user.id}`
    };
    return this.httpClient.post<UserResponse>(environment.identifyRequestURL+"/users", body);
  }

  login(): void {
    this.postLoggedInUser()
    .subscribe(
      (userResponse: UserResponse) => {
        console.log("logged in");
        localStorage.setItem("auth",userResponse.data.token);
        this.router.navigate(['mainpage']);
      },
      err => {
        console.error("Could not login:" + err);
      }
    );
    this.triedLogIn = true;
  }

}
