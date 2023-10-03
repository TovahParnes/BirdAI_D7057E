import { SocialAuthService } from '@abacritt/angularx-social-login';
import { Component } from '@angular/core';
import {AuthGuardService} from './services/auth-guard.service';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})

export class AppComponent {
  isDark = true;

  switchDarkmodeSetting(): void {
    document.documentElement.style.display = 'none';
    document.documentElement.setAttribute(
        "data-color-scheme",
        this.isDark  ? "dark" : "light"
    );
    document.body.clientWidth;
    document.documentElement.style.display = '';
    this.isDark = true;
  }

  user:any;
  loggedIn: boolean = false;

  constructor(private authService: SocialAuthService) { }

  ngOnInit() {
    this.authService.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null);
      localStorage.setItem('id_token', user.idToken);
    });
  }

}