import { Component } from '@angular/core';
import { ArtService } from '../services/art.service';

@Component({
  selector: 'app-add-art',
  templateUrl: './add-art.component.html',
  styleUrls: ['./add-art.component.css']
})
export class AddArtComponent {

  // Object bound to the form
  artwork = {
    artName: '',
    artist: '',
    description: '',
    photoUrl:''
  };

  successMessage = '';
  errorMessage = '';
  loading = false;

  constructor(private artService: ArtService) { }

  // Called when form is submitted
  onSubmit() {
    // Simple validation
    if (!this.artwork.artName || !this.artwork.artist) {
      this.errorMessage = 'Art Name and Artist are required.';
      this.successMessage = '';
      return;
    }

    this.loading = true;
    this.artService.addArtwork(this.artwork).subscribe({
      next: (res) => {
        this.successMessage = 'Artwork added successfully!';
        this.errorMessage = '';
        this.loading = false;

        // Reset form
        this.artwork = { artName: '', artist: '', description: '',  photoUrl:''};
      },
      error: (err) => {
        console.error('Add artwork error:', err);
        this.errorMessage = 'Failed to add artwork. Check console.';
        this.successMessage = '';
        this.loading = false;
      }
    });
  }
}
