import {Component} from '@angular/core';
import {SocialAuthService} from 'angularx-social-login';
import {Router} from '@angular/router';
import {Directive} from '@angular/core'

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})

export class MainPageComponent {

  constructor(private router: Router,
              public socialAuthServive: SocialAuthService) {
  }

  logout(): void {
    this.socialAuthServive.signOut().then(() => this.router.navigate(['login']));
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

}

// export class ImageUploadComponent {
//   imageUrls: string[] = [];

//   onDragOver(event: DragEvent): void {
//     event.preventDefault();
//     event.stopPropagation();
//   }

//   onDragLeave(event: DragEvent): void {
//     event.preventDefault();
//     event.stopPropagation();
//   }

//   onDrop(event: DragEvent): void {
//     event.preventDefault();
//     event.stopPropagation();
//     const files = event.dataTransfer?.files;
//     if (files) {this.processFiles(files);
//     }
//   }

//   onFileSelected(event: Event): void {
//     const inputElement = event.target as HTMLInputElement;
//     const files = inputElement.files;
//     this.processFiles(files);
//     inputElement.value = ''; // Reset the input value to allow re-uploading the same file
//   }

//   private processFiles(files: FileList | null): void {
//     if (!files) return;

//     for (let i = 0; i < files.length; i++) {
//       const file = files[i];
//       if (file.type.startsWith('image/')) {
//         const imageUrl = URL.createObjectURL(file);
//         this.imageUrls.push(imageUrl);
//       }
//     }
//   }
// }