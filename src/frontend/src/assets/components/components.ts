

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
