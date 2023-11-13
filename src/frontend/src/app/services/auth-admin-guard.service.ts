import { SocialAuthService, SocialUser } from '@abacritt/angularx-social-login';
import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Router, RouterStateSnapshot } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { AdminResponse } from 'src/assets/components/components';

@Injectable({
    providedIn: 'root' 
})

export class AuthGuardAdminService {
  public user: SocialUser | undefined;
  public loggedIn = false;
  private currentAdmin = localStorage.getItem("currentAdmin")

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

  async getCurrentAdmin(){
    const authKey = localStorage.getItem("auth");
    if(authKey){
      (await this.sendGetCurrentAdmin(authKey)).subscribe(
        (response: AdminResponse) => {
          this.currentAdmin = response.data._id;
        }
      )
    }
  }

  async sendGetCurrentAdmin(token:string){
    const header = {
      'Authorization': `Bearer ${token}`
    };
    return this.http.get<AdminResponse>(environment.identifyRequestURL+"/admins/me",{ headers: header });
  }

  async canActivate(state: RouterStateSnapshot){
    const currentUser = localStorage.getItem('userId');
    await this.getCurrentAdmin();
    if (this.currentAdmin == currentUser){
        console.log("ok")
    }else {
        console.log("notok")
        this.router.navigate(['login'], { queryParams: { returnUrl: state.url }});
    }
  }
}