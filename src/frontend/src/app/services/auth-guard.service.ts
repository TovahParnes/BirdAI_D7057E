import { SocialAuthService, SocialUser } from '@abacritt/angularx-social-login';
import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Router, RouterStateSnapshot } from '@angular/router';

@Injectable({
    providedIn: 'root'
})

export class AuthGuardService {
  public user: SocialUser | undefined;
  public loggedIn = false;

  constructor(
    private router: Router,
    private authService: SocialAuthService) {
    this.authService.authState.subscribe(async (user) => {
      this.user = user;
      this.loggedIn = (user != null);
      localStorage.setItem('id_token', user.idToken);
    });
  }

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot){
    const logged = this.loggedIn;
    const destination: string = state.url;

    if (!this.loggedIn) {
        console.log(state.url);
        this.router.navigate(['login'], { queryParams: { returnUrl: state.url }});
    }
}
}
