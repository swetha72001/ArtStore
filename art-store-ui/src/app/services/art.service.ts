import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root' // makes this service available globally across the app
})
export class ArtService {

  private baseUrl = 'http://localhost:8090'; // Go backend base URL

  constructor(private http: HttpClient) { } //to send hhtp call

  // To add (insert) artwork
  addArtwork(art: any): Observable<any> {
    console.log("art>>>>>>>",art)
    return this.http.post(`${this.baseUrl}/artworks`, art);
  }

  // To get list of artworks
  getArtworks(): Observable<any> {
    return this.http.get(`${this.baseUrl}/listartworks`);
  }
}
