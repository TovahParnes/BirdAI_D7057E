export interface AnalyzeResponse {
    "timestamp": string;
    "data": BirdData[];
}

interface BirdData {
    "aiBird": {
        "name": string;
        "accuracy": Number;
    };
    "birdId": string;
    "description":string;
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
    "token": string,
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
    "accuracy": Number,
    "bird": Bird,
    "comment": string,
    "createdAt": string,
    "location": string,
    "user": ListUser,
    "userMedia": Media
}

interface Bird {
    "Id": string,
    "Name": string,
    "Description": string,
    "Sound": string,
    "Species": boolean,
    "Image": string
}

export interface PostData {
    "accuracy": Number,
    "birdId": string,
    "comment": string,
    "location": string,
    "imageId": string,
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
      "accuracy": Number;
    }[]
  }

 export interface AnalyzedBird {
    "title": string;
    "image": string;
    "accuracy": Number;
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

export interface getFoundBirds {
    "data": PostData[],
    "timestamp": string
}

export interface SoundSegment {
  "startTime": number,
  "data": string
}
