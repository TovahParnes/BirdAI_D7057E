import {Component, ElementRef, ViewChild} from '@angular/core';
import WaveSurfer from "wavesurfer.js";
import RegionsPlugin from "wavesurfer.js/dist/plugins/regions";

@Component({
  selector: 'app-sound-editor',
  templateUrl: './sound-editor.component.html',
  styleUrls: ['./sound-editor.component.css']
})
export class SoundEditorComponent {

  @ViewChild('waveform', { static: false }) waveform!: ElementRef

  private wavesurfer!: WaveSurfer

  // @ts-ignore
  private wsRegions: WaveSurfer.Plugin & { regions: WaveSurfer.Region[] }

  // @ts-ignore
  private activeRegion: WaveSurfer.Region

  // Sound editor controls
  private loopRegion: boolean = true
  public isPaused: boolean = true
  public soundLoaded: boolean = false

  // Maximum allowed audio length to be loaded into frontend for visualization
  private maxAudioLength: number = 120

  // Allowed audio types
  private allowedTypes: string[] = ['audio/wav', 'audio/mpeg']

  /**
   * Initializes the Wavesurfer instance with the provided configuration.
   * Sets up event listeners for the 'ready' event, 'region-in' event, 'region-out' event,
   * 'region-clicked' event, and 'interaction' event to handle various interactions with the audio waveform.
   */
  private initWavesurfer(): void {
    this.soundLoaded = true

    this.wavesurfer = WaveSurfer.create({
      container: this.waveform.nativeElement,
      waveColor: 'rgb(200, 0, 200)',
      progressColor: 'rgb(100, 0, 100)',
    });

    this.wavesurfer.on('ready', () : void => {

      // @ts-ignore
      this.wsRegions.on('region-in', (region) : void  => {
        this.activeRegion = region
        this.isPaused = false
      });

      // @ts-ignore
      this.wsRegions.on('region-out', (region) : void  => {
        if (this.activeRegion === region) {
          if (this.loopRegion) {
            region.play()
          } else {
            this.activeRegion = null
            this.isPaused = true
          }
        }
      });

      // @ts-ignore
      this.wsRegions.on('region-clicked', (region, e) : void => {
        e.stopPropagation() // prevent triggering a click on the waveform
        this.activeRegion = region
        this.isPaused = false
        region.play()
        region.setOptions({ color: this.randomColor() })
      });

      // Reset the active region when the user clicks anywhere in the waveform
      this.wavesurfer.on('interaction', () : void => {
        this.activeRegion = null
      })
    });
  }

  /**
   * Toggles the playback state of the audio track.
   * If the track is playing, it will be paused, and if it's paused, it will resume playing.
   * If there is an active region, playing will resume from the active region's start point.
   * The method updates the `isPaused` flag accordingly.
   */
  public togglePauseTrack() : void {
    if (!this.isPaused && this.wavesurfer.isPlaying()) this.wavesurfer.pause()
    if (this.isPaused && this.activeRegion) this.activeRegion.play()
    this.isPaused = !this.isPaused
  }

  /**
   * Event handler for file input change. Resets the existing Wavesurfer instance (if any),
   * initializes a new Wavesurfer instance, and processes the selected sound file.
   *
   * @param {any} event - The input change event containing the selected files.
   */
  onFileChange(event: any): void {
    if (this.wavesurfer) this.resetComponent()  // If existing wavesurfer, reset it.

    this.initWavesurfer()
    this.initRegionsPlugin()

    if (event.target.files.length > 0) {
      const file: File = event.target.files[0]

      if (this.isFileTypeAllowed(file)) {
        this.readSoundFile(file)
        this.checkFileDuration(file)
      } else {
        console.error('Invalid file type. Please upload a .wav or .mp3 file.')
      }
    }
  }

  /**
   * Checks the duration of the provided sound file and logs an error if it exceeds the maximum allowed duration.
   *
   * @param {File} file - The sound file to check for duration.
   */
  checkFileDuration(file: File): void {
    const audio: HTMLAudioElement = new Audio()
    const reader: FileReader = new FileReader()

    reader.onload = (e: ProgressEvent<FileReader>) : void => {
      if (typeof e.target?.result === 'string') {
        audio.src = e.target?.result
        audio.onloadedmetadata = () : void => {
          const duration: number = audio.duration
          if (duration > this.maxAudioLength) console.error('File duration is greater than', this.maxAudioLength, 'seconds. Please upload a shorter file.')
        };
      }
    };

    reader.readAsDataURL(file)
  }

  /**
   * Checks if the file type is allowed for sound processing.
   *
   * @param {File} file - The File object representing the sound file.
   * @returns {boolean} True if the file type is allowed, false otherwise.
   */
  private isFileTypeAllowed(file: File): boolean {
    return this.allowedTypes.includes(file.type)
  }

  /**
   * Reads the content of a sound file using FileReader and displays the sound data.
   *
   * @param {File} file - The File object representing the sound file.
   */
  private readSoundFile(file: File): void {
    const reader : FileReader = new FileReader()

    reader.onloadend = () : void => {
      const soundData : string = reader.result as string
      this.displaySound(soundData)
    };

    reader.readAsDataURL(file)
  }

  /**
   * Displays the sound data by loading it into the Wavesurfer instance and adding a default region.
   *
   * @param {string} soundData - The sound data to be displayed.
   */
  private displaySound(soundData: string): void {
    if (this.wavesurfer) this.wavesurfer.load(soundData)
    this.addRegion()
  }

  /**
   * Generates a random number within the specified range.
   *
   * @param {number} min - The minimum value of the range.
   * @param {number} max - The maximum value of the range.
   * @returns {number} A random number within the specified range.
   */
  private random(min: number, max: number): number {
    return Math.random() * (max - min) + min
  }

  /**
   * Generates a random RGBA color with a 50% opacity.
   *
   * @returns {string} The randomly generated RGBA color.
   */
  private randomColor = () : string => `rgba(${this.random(0, 255)}, ${this.random(0, 255)}, ${this.random(0, 255)}, 0.5)`

  /**
   * Adds a region to the Wavesurfer instance when the audio is decoded.
   * The added region spans from 0 to 10 seconds by default and has a randomly generated color.
   * The region is non-resizable.
   */
  private addRegion() : void {
    this.wavesurfer.on('decode', () : void => {
      this.wsRegions.addRegion({
        start: 0,
        end: 10,
        content: 'Choose 10 seconds that fits the best, Drag me',
        color: this.randomColor(),
        resize: false,
      })
    })
  }

  /**
   * Initializes the RegionsPlugin for the Wavesurfer instance.
   * This method registers the plugin and assigns it to the wsRegions property.
   */
  private initRegionsPlugin(): void {
    this.wsRegions = this.wavesurfer.registerPlugin(RegionsPlugin.create())
  }

  /**
   * Resets the state of the SoundEditorComponent.
   * Clears the active region, destroys the Wavesurfer instance, and sets the regions plugin to null.
   */
  resetComponent() : void {
    this.soundLoaded = false
    this.wavesurfer.destroy()
    this.wsRegions = null
  }

  public saveSection() : void {
    console.log("This should be the audio data we use for saving: ", this.wsRegions.regions[0])
    /* TODO: Implement service call to backend, pass start time and data. */
  }
}
