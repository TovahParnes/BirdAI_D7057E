import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

export class WikiSummary {
    type?:          string;
    title?:         string;
    displaytitle?:  string;
    //namespace?:     Namespace;
    wikibase_item?: string;
    titles?:        Titles;
    pageid?:        number;
    //thumbnail?:     Originalimage;
    originalimage?: Originalimage;
    lang?:          string;
    dir?:           string;
    revision?:      string;
    tid?:           string;
    timestamp?:     Date;
    description?:   string;
    //content_urls?:  ContentUrls;
    //api_urls?:      APIUrls;
    extract?:       string;
    extract_html?:  string;
    note?: string;
  }

  export class WikiPageSegment{
    sourcelanguage?: string;
    title?: string;
    revision?: string;
    segmentedContent?: string;
  }

  export interface WikiImages{
    items: WikiImageItems[];
  }

  interface WikiImageItems{
    title: string;
    srcset: WikiImagesSrcSet[];
  }

  interface WikiImagesSrcSet {
      src: string,
      scale: string,
  }

  interface Originalimage{
    height?: number,
    source?: string,
    width?: number,
  }

  interface Titles{
    canoncical?: string,
    display?: string,
    normalized?: string,
  }

  @Injectable({
    providedIn: 'root' 
})

export class WikirestService {
    constructor( private http:  HttpClient) { }
  
    getWiki(title: string) {
      const tempTitle = title.replace(' ', '_') + '?redirect=true';
      const baseUrl = 'https://en.wikipedia.org/api/rest_v1/page/summary/';
      return this.http.get<WikiSummary>(baseUrl+tempTitle);
    }

    getWikiPage(title:string){
        const tempTitle = title.replace(' ', '_') + '?redirect=true';
        const baseUrl = 'https://en.wikipedia.org/api/rest_v1/page/html/';
        return this.http.get<string>(baseUrl+tempTitle);
    }

    getWikiImages(title:string){
      const tempTitle = title.replace(' ', '_') + '?redirect=true';
      const baseUrl = 'https://en.wikipedia.org/api/rest_v1/page/media-list/';
      return this.http.get<WikiImages>(baseUrl+tempTitle);
    }
  }