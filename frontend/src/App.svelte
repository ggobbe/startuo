<script lang="ts">
  import logo from './assets/images/back.png'
  import {DownloadManifest} from '../wailsjs/go/main/App.js'

  let resultText: string = ""

  function downloadManifest(): void {
    DownloadManifest().then(result => resultText = result)
  }
</script>

<main>
  <img alt="logo" id="logo" src="{logo}">
  <div class="input-box" id="input">
    <button class="btn" on:click={downloadManifest}>Download Manifest</button>
  </div>
  <hr>
  <pre class="result" id="result">{resultText}</pre>
</main>

<style>

  #logo {
    display: block;
    width: 100%;
    margin: auto;
    padding: 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: auto;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
