import {Component} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import {AppComponent} from '../app.component';


@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.css']
})

export class AboutComponent {

  constructor(
    private router: Router,
    public socialAuthService: SocialAuthService,
    public mainApp: AppComponent,
  ){}
}
