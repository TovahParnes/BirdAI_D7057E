import {Component} from '@angular/core';
import {SocialAuthService, GoogleLoginProvider} from 'angularx-social-login';
import {Router} from '@angular/router';
//import {Directive} from '@angular/core'
import { HttpClient } from '@angular/common/http';
//import { Observable } from 'rxjs';


interface ApiResponse {
  data: {  
    id : string;
    authId: string;
    createdAt: string;
    username: string;
  }[];
  message: string;
  success: boolean;
}

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css'],
})

export class MainPageComponent {

  responseData: ApiResponse | null = null;

  constructor(private router: Router,
              public socialAuthService: SocialAuthService,
              private http: HttpClient,
              ) {
  }

  logout(): void {
    this.socialAuthService.signOut().then(() => this.router.navigate(['login']));
  }

  moveToLibrary(): void {
    this.router.navigate(['library']);
  }

  imageUrls: string[] = [];

  onDragOver(event: DragEvent): void {
    event.preventDefault();
    event.stopPropagation();
  }

  onDragLeave(event: DragEvent): void {
    event.preventDefault();
    event.stopPropagation();
  }

  onDrop(event: DragEvent): void {
    event.preventDefault();
    event.stopPropagation();
    const files = event.dataTransfer?.files;
    if (files) {this.processFiles(files);
    }
  }

  onFileSelected(event: Event): void {
    const inputElement = event.target as HTMLInputElement;
    const files = inputElement.files;
    this.processFiles(files);
    inputElement.value = ''; // Reset the input value to allow re-uploading the same file
  }

  private processFiles(files: FileList | null): void {
    if (!files) return;

    for (let i = 0; i < files.length; i++) {
      const file = files[i];
      if (file.type.startsWith('image/')) {
        const imageUrl = URL.createObjectURL(file);
        this.imageUrls.push(imageUrl);
      }
    }
  }

  ngOnInit(): void {
    // Make an HTTP GET request to the Swagger service's API
    this.http.get<ApiResponse>('http://localhost:4000/swagger/index.html').subscribe(data => {
      this.responseData = data;
      console.log(data);
    });
  }

}

