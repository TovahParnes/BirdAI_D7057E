import {Component, OnInit, ViewChild} from '@angular/core';
import {SocialAuthService, SocialUser} from '@abacritt/angularx-social-login';
import {Router} from '@angular/router';
import {AppComponent} from '../app.component';
import { FormBuilder, FormGroup, FormControl } from '@angular/forms';
import { MatStepper } from '@angular/material/stepper';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css'],
})

export class MainPageComponent implements OnInit {

  user: SocialUser = new SocialUser;
  loggedIn: boolean = false;
  isLinear = false;
  form!: FormGroup;
  selectedImage: any;

  constructor(
    private router: Router, 
    public mainApp: AppComponent,
    public socialAuthService: SocialAuthService,
    private formBuilder: FormBuilder) {
  }

  ngOnInit() {
      this.socialAuthService.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null);
    }),

    this.form = this.formBuilder.group({
      option: new FormControl(), // Initialize with a default value
    });
  }

  @ViewChild('stepper') stepper!: MatStepper;

  ngAfterViewInit() {
    this.stepper.selectedIndex = -1;
  }

  logout(): void {
    this.socialAuthService.signOut().then(() => this.router.navigate(['login']));
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

  toggleTheme(): void {
    this.mainApp.switchDarkmodeSetting();
  }

  onFileSelected(event: any) {
    const file = event.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = (e) => {
        this.selectedImage = reader.result;
      };
    }
  }

  onSubmit() {
    // Handle form submission here, e.g., send the image to a server
    // You can access the selected image using this.selectedImage
    console.log('Image submitted:', this.selectedImage);

    // Reset the form or perform any other necessary actions
    this.form.reset();
    this.selectedImage = null;
  }
}
