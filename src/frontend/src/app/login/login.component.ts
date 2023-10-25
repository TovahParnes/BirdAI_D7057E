
import {Component} from '@angular/core';
import {GoogleLoginProvider, SocialAuthService} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { UserResponse } from 'src/assets/components/components';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  constructor(
    public mainApp: AppComponent,
    private router: Router,
    private httpClient: HttpClient,
    private socialAuthService: SocialAuthService) {
  }
  
  public triedLogIn = false;

  postLoggedInUser(): Observable<UserResponse> {
    const body = {
      'username': `${this.mainApp.user.name}`,'authId': `${this.mainApp.user.id}`
    };
    return this.httpClient.post<UserResponse>(environment.identifyRequestURL+"/users", body);
  }

  login(): void {
    console.log(this.mainApp.user);
    this.postLoggedInUser()
    .subscribe(
      (userResponse: UserResponse) => {
        console.log("logged in");
        console.log(userResponse);
        // localStorage.setItem("auth",userResponse.data.authId);
        // console.log(localStorage.getItem("auth"));
        this.router.navigate(['mainpage']);
      },
      err => {
        console.error("Could not login:" + err);
      }
    );
    this.triedLogIn = true;
  }
}
