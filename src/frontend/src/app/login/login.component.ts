import {Component} from '@angular/core';
import {GoogleLoginProvider, SocialAuthService} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import { AppComponent } from '../app.component';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  constructor(private router: Router,
              public mainApp: AppComponent,
              private socialAuthService: SocialAuthService) {
  }
  
  navigateToHome(): void {
    this.router.navigate(['mainpage']);
  }
}
