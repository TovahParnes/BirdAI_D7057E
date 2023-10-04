import {Component} from '@angular/core';
import {GoogleLoginProvider, SocialAuthService} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import { Observable } from 'rxjs';

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
  
  postLoggedInUser(token: string): Observable<any> {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`,
      'authId': `${this.mainApp.user.id}`
    })
    return this.httpClient.post("localhost:4000/users", { headers: headers })
  }


  login(): void {
    this.postLoggedInUser("ss").subscribe(
      error => console.log("Could not login"),
      () => this.router.navigate(['mainpage'])
    );
  }
}
