import { SocialAuthService, SocialUser } from '@abacritt/angularx-social-login';
import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Router, RouterStateSnapshot } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { AdminResponse } from 'src/assets/components/components';

@Injectable({
    providedIn: 'root' 
})

export class AuthGuardService {
  public user: SocialUser | undefined;
  public loggedIn = false;
  private currentAdmin?: AdminResponse

  constructor(
    private router: Router, 
    private http: HttpClient,
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
        this.router.navigate(['login'], { queryParams: { returnUrl: state.url }});
    }
}
}