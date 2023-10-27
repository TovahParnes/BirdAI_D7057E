

export interface AnalyzeResponse {
    timestamp: string;
    data: BirdData;
}

interface BirdData {
    accuracy: string;
    name: string;
    picture: Media;
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
        "data": listOutputData,
        "timestamp": string
}

export interface listOutputData {
    "bird": Bird,
    "createdAt": string,
    "_id": string,
    "location": string,
    "user": ListUser,
    "usermedia": Media
}

interface Bird {
    "description": string,
    "id": string,
    "image": Media,
    "name": string,
    "sound": Media,
    "filetype": string,
    "_id":string
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