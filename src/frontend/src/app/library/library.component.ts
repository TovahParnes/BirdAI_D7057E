import {Component} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider} from 'angularx-social-login';
import {Router} from '@angular/router';

@Component({
  selector: 'app-library',
  templateUrl: './library.component.html',
  styleUrls: ['./library.component.css']
})

export class LibraryComponent {

  constructor(private router: Router,
    public socialAuthService: SocialAuthService) {

  }

  moveToHome(){
    this.router.navigate(['mainpage']);
  }

  moveToSpeciesPage(){
    this.router.navigate(['species-page']);
  }
}
