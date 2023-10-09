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
    private http: HttpClient,
    private socialAuthService: SocialAuthService) {
  }

  public triedLogIn = false;

  postLoggedInUser(token: string): Observable<any> {
    var header = {
      headers: new HttpHeaders()
        .set('Authorization',  `Bearer ${token}`)
    }
    //console.log(header);
    const body = {
      'username': `${this.mainApp.user.name}`,'authId': `${this.mainApp.user.id}`
    };
    // return this.httpClient.post<any>(environment.loginURL, body, { headers: header });
    return this.http.post<any>(environment.loginURL, body, header);
  }


  login(): void {
    console.log(this.mainApp.user);
    //this.router.navigate(['mainpage']); // TODO-REMOVE

    this.postLoggedInUser(environment.secret)
    .subscribe(
      (value: any) => {
        console.log("logged in");
        console.log(value);
        this.router.navigate(['mainpage']);
      },
      err => {
        console.error("Could not login:" + err);
      }
    );
    this.triedLogIn = true;
  }
}
