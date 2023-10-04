import {Component} from '@angular/core';
import {GoogleLoginProvider, SocialAuthService} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

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

  postLoggedInUser(token: string): Observable<any> {
    const header = {
      'Authorization': `Bearer ${token}`
    };
    const body = {
      'username': `${this.mainApp.user.name}`,'authId': `${this.mainApp.user.id}`
    };
    return this.httpClient.post<any>(environment.loginURL, body, { headers: header });
  }


  login(): void {
    console.log(this.mainApp.user);
    this.router.navigate(['mainpage']); // TODO-REMOVE

    // this.postLoggedInUser("environment.secret")
    // .subscribe(
    //   () => {
    //     console.log("logged in");
    //     this.router.navigate(['mainpage']);
    //   },
    //   err => {
    //     console.error("Could not login:" + err);
    //   }
    // );
    this.triedLogIn = true;
  }
}
