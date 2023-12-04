package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"strings"

	"github.com/hajimehoshi/go-mp3"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// RequestAnalyze Connects to AI and retrieves the name and accuracy of the picture
func (c *Controller) RequestAnalyze(mediaData, endpoint string) (models.AIList, error) {
	if os.Getenv("USE_AI") == "true" {
		postBody, _ := json.Marshal(map[string]string{
			"media": mediaData,
		})
		responseBody := bytes.NewBuffer(postBody)
		resp, err := http.Post(os.Getenv("AI_URI")+"/"+endpoint, "application/json", responseBody)
		if err != nil {
			return models.AIList{}, models.Err{
				StatusCode:  http.StatusServiceUnavailable,
				StatusName:  http.StatusText(http.StatusServiceUnavailable),
				Message:     "Internal AI error",
				Description: err.Error(),
			}
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var aiBird models.AIList
		err = json.Unmarshal(body, &aiBird)
		if err != nil {
			var e models.AIError
			err = json.Unmarshal(body, &e)
			if err != nil {
				return models.AIList{}, models.Err{
					StatusCode:  http.StatusInternalServerError,
					StatusName:  http.StatusText(http.StatusInternalServerError),
					Message:     "Unknown Internal AI error",
					Description: err.Error(),
				}
			}
			return models.AIList{}, models.Err{
				StatusCode:  http.StatusInternalServerError,
				StatusName:  http.StatusText(http.StatusInternalServerError),
				Message:     "Internal AI error",
				Description: e.Error,
			}
		}
		return aiBird, nil
	}
	return models.AIList{Birds: []models.AIBird{
		{
			Name:     "EURASIAN MAGPIE",
			Accuracy: 1,
		},
	}}, nil
}

func (c *Controller) AiListToResponse(aiList models.AIList) []models.AnalyzeResponse {
	response := []models.AnalyzeResponse{}
	for _, ai := range aiList.Birds {
		bird := c.CGetBirdByName(strings.ToUpper(ai.Name))
		if utils.IsTypeError(bird) {
			continue
		}
		response = append(response, models.AnalyzeResponse{
			AiBird: models.AIBird{
				Name:     ai.Name,
				Accuracy: ai.Accuracy,
			},
			BirdId:      bird.Data.(*models.BirdOutput).Id,
			Description: bird.Data.(*models.BirdOutput).Description,
		})
	}
	return response
}

// IsWAV checks if the provided data represents a WAV file
func (c *Controller) IsWAV(data []byte) bool {
	return len(data) > 11 && string(data[8:12]) == "WAVE"
}

// IsMP3 checks if the provided data represents an MP3 file
func (c *Controller) IsMP3(data []byte) bool {
	return len(data) > 2 && string(data[:3]) == "ID3"
}

// DecodeWAV decodes WAV audio data
func (c *Controller) DecodeWAV(data []byte) (*audio.IntBuffer, error) {
	reader := bytes.NewReader(data)
	dec := wav.NewDecoder(reader)
	if dec.Err() != nil {
		return nil, dec.Err()
	}
	return dec.FullPCMBuffer()
}

// DecodeMP3 decodes MP3 audio data
func (c *Controller) DecodeMP3(data []byte) (*audio.IntBuffer, error) {
	decoder, err := mp3.NewDecoder(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 8192)
	audioData := audio.IntBuffer{}

	for {
		n, err := decoder.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		// Convert 16-bit little-endian samples to int and append to the audio data
		for i := 0; i < n; i += 2 {
			sample := int(int16(buf[i]) | int16(buf[i+1])<<8)
			audioData.Data = append(audioData.Data, sample)
		}
	}

	audioData.Format = &audio.Format{
		SampleRate:  decoder.SampleRate(),
		NumChannels: 2, // MP3 is always stereo (2 channels)
	}
	audioData.SourceBitDepth = 16

	return &audioData, nil
}

// TrimAudioData trims the audio data to the specified duration
func (c *Controller) TrimAudioData(audioData *audio.IntBuffer, targetDuration, start time.Duration) *audio.IntBuffer {
	// Calculate the number of samples corresponding to the target duration
	targetSamples := int(float64(targetDuration) * float64(audioData.Format.SampleRate) / float64(time.Second))
	startSamples := int(float64(start) * float64(audioData.Format.SampleRate) / float64(time.Second))

	// Trim audio data to the target number of samples
	if len(audioData.Data) > targetSamples+startSamples {
		audioData.Data = audioData.Data[startSamples : targetSamples+startSamples]
	} else if len(audioData.Data) > targetSamples {
		audioData.Data = audioData.Data[len(audioData.Data)-targetSamples : len(audioData.Data)]
	} else {
		newData := []int{}
		for len(newData) < targetSamples {
			newData = append(newData, audioData.Data...)
		}
		audioData.Data = newData[:targetSamples]
	}

	return audioData
}

// EncodeAudioData encodes the audio data to the specified format and returns the base64 representation
func (c *Controller) EncodeAudioData(audioData *audio.IntBuffer) string {
	var encodedData []byte
	var err error

	// Create a temporary WAV file
	tempFile, err := ioutil.TempFile("", "temp.wav")
	if err != nil {
		log.Fatal("Error creating temporary WAV file:", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Create a WAV encoder with the file
	encoder := wav.NewEncoder(tempFile, audioData.Format.SampleRate, 16, audioData.Format.NumChannels, 1)

	// Write the audio data to the WAV encoder
	if err := encoder.Write(audioData); err != nil {
		log.Fatal("Error encoding trimmed WAV:", err)
	}

	// Close the encoder to flush any remaining data
	if err := encoder.Close(); err != nil {
		log.Fatal("Error closing WAV encoder:", err)
	}

	// Read the encoded WAV data from the temporary file
	encodedData, err = ioutil.ReadFile(tempFile.Name())
	if err != nil {
		log.Fatal("Error reading temporary WAV file:", err)
	}

	if err != nil {
		log.Fatal("Error encoding audio:", err)
	}

	return base64.StdEncoding.EncodeToString(encodedData)
}
