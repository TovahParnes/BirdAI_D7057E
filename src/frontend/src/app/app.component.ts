import {Component} from '@angular/core';
import {SocialAuthService} from '@abacritt/angularx-social-login';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})

export class AppComponent {
  user:any;
  loggedIn: boolean = false;

  constructor(
    public authService: SocialAuthService) {
  }

  ngOnInit() {
    this.authService.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null);
      localStorage.setItem('id_token', user.idToken);
    });
  }
}
