import { SocialAuthService, SocialUser } from '@abacritt/angularx-social-login';
import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Router, RouterStateSnapshot } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { AdminResponse } from 'src/assets/components/components';
import { UserResponse } from 'src/assets/components/components';

@Injectable({
    providedIn: 'root'
})

export class AuthGuardService {
  public user: SocialUser | undefined;
  public loggedIn = false;
  private currentAdmin?: AdminResponse;

  constructor(
    private router: Router,
    private authService: SocialAuthService,
    private http: HttpClient) {
    this.authService.authState.subscribe(async (user) => {
      this.user = user;
      this.loggedIn = (user != null);
      localStorage.setItem("userPhoto", user.photoUrl);
      localStorage.setItem('id_token', user.idToken);
      localStorage.setItem("username",user.name);
    });
  }

  getCurrentUser(token: string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    return this.http.get<UserResponse>(environment.identifyRequestURL+"/users/me",{ headers: header });
  }

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot){
    const authKey = localStorage.getItem("auth");

    if(authKey){
      this.getCurrentUser(authKey).subscribe(
        (response: UserResponse) => {
          
      },err => { 
        console.log(state.url);
        this.router.navigate(['first-page'], { queryParams: { returnUrl: state.url }});
        console.error("Not Logged In:" + err); 
      }
    )}else{
      this.router.navigate(['first-page'], { queryParams: { returnUrl: state.url }});
    }
  }
}
