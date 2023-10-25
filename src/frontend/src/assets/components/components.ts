

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
    "data": User,
    "timestamp": string,
}

export interface User {
    "_id": string,
    "username": string,
    "authId": string,
    "createdAt": string,
    "Active": boolean
}
