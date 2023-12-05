import { SocialAuthService } from "@abacritt/angularx-social-login";
import { Component, Input, booleanAttribute } from "@angular/core";
import { Router } from "@angular/router";
import { AppComponent } from "../app.component";


@Component({
    selector: 'app-navbar',
    templateUrl: './navbar.component.html',
    styleUrls: ['./navbar.component.css'],
    template: '<navbar firstSelected="true" secondSelected="false" thirdSelected="false"></navbar>'
})
  
export class NavbarComponent {

    constructor(
        private router: Router,
        public socialAuthService: SocialAuthService,
        public mainApp: AppComponent) {
    }

    userPhoto = localStorage.getItem("userPhoto");
        
    @Input({ transform: booleanAttribute }) firstSelected!: boolean;
    @Input({ transform: booleanAttribute }) secondSelected!: boolean;
    @Input({ transform: booleanAttribute }) thirdSelected!: boolean;


    logout(): void {
        this.mainApp.authService.signOut();
        localStorage.setItem("auth","-1");
        localStorage.setItem("username","");
        localStorage.setItem("userPhoto","");
        this.router.navigate(["first-page"]);
    }
    
    navigateToHome(): void {
        this.router.navigate(['mainpage']);
    }
    
    navigateToTakenImages(): void {
        this.router.navigate(['takenImages']);
    }
    
    navigateToLibrary(): void {
        this.router.navigate(['library']);
    }
    
    navigateToProfilePage(): void {
        this.router.navigate(['profile']);
    }
    
    getLoggedIn(){
        return localStorage.getItem("loggedIn");
    }
} 
  