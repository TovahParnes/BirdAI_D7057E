

export interface AnalyzeResponse {
    "timestamp": string;
    "data": BirdData[];
}

interface BirdData {
    "aiBird": { 
        "name": string; 
        "accuracy": string;
    };
    "birdId": string;
    "userMedia": Media;
}

interface Media {
    "_id": string;
    "data": string;
    "fileType": string;
}

export interface UserResponse {
    "data": LoginUser,
    "timestamp": string,
}

export interface LoginUser {
    "_id": string,
    "active": boolean
    "authId": string,
    "createdAt": string,
    "username": string,
}

export interface ListUser {
    "_id": string,
    "username": string,
    "createdAt": string,
    "Active": boolean
}

export interface listOutput {
        "timestamp": string,
        "data": listOutputData[]
}

export interface listOutputData {
    "_id": string,
    "user": ListUser,
    "bird": Bird,
    "createdAt": string,
    "location": string,
    "userMedia": Media
}

interface Bird {
    "Id": string,
    "Name": string,
    "Description": string,
    "Image": Media,
    "Sound": Media,
    "species": Boolean
}

export interface PostData {
    "birdId": string,
    "imageId": string,
    "location": string,
    "soundId": string
}

export interface Post {
    "data": PostData,
    "timestamp": string
}

export interface UserBirdList {
    birds:{   
      "title": string;
      "image": string;
      "accuracy": string;
    }[]
  }
  
 export interface AnalyzedBird {
    "title": string;
    "image": string;
    "accuracy": string;
  }

export interface DeleteResponse {
    "timestamp": string;
    "data": string;
}

export interface UpdateResponse {
    "timestamp": string,
    "data": {
        "Id": string,
        "UserId": string,
        "BirdId": string,
        "CreatedAt": string,
        "Location": string,
        "MediaId": string
    }
}
export interface getAllBirdsResponse {
    "timestamp": string;
    "data": Bird[];
}

export interface AdminResponse {
    "timestamp": string;
    "data": AdminData;
}

interface AdminData {
    "_id": string;
    "user": ListUser;
    "access": string;
}