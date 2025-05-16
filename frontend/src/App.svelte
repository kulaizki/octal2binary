<script lang="ts">
  import { OctalToBinary } from '../wailsjs/go/main/App'; 

  let octalInput: string = "";
  let binaryOutput: string = "Enter an octal string (e.g., 1031, 00507) and click 'Convert'";
  let errorText: string = ""; 

  async function convertOctal(): Promise<void> {
    if (typeof octalInput !== 'string') {
      errorText = "Invalid input type.";
      binaryOutput = "";
      return;
    }
    
    errorText = ""; 

    try {
      const result = await OctalToBinary(octalInput.trim());
      binaryOutput = result;
    } catch (err: any) {
      binaryOutput = "";
      if (typeof err === 'string') {
        errorText = err;
      } else if (err && typeof err.message === 'string') {
        errorText = err.message;
      } else {
        errorText = "An unknown error occurred during conversion.";
      }
      console.error("Conversion error:", err);
    }
  }
</script>

<main>
  <div class="container">
    <div class="app-header">
      <h1>Octal to Binary Converter</h1>
      <p class="subtitle">Converts an octal string to its binary representation</p>
    </div>
    
    <div class="card">
      <div class="input-section">
        <label for="octal-input">Octal Input</label>
        <textarea 
          id="octal-input"
          bind:value={octalInput} 
          placeholder="e.g., 1031, 00507, 0"
          rows="4"
          on:input={() => { errorText = ''; }}
        ></textarea>
      </div>
      
      <button class="hover:cursor-pointer reverse-btn" on:click={convertOctal}>
        <span>Convert to Binary</span>
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 12H4M12 20l-8-8 8-8"/></svg> <!-- Changed icon to a conversion-like arrow -->
      </button>
      
      <div class="result-section">
        <label for="binary-output">Binary Output</label>
        <div id="binary-output" class="result-box" class:error={errorText !== ''}>
          {#if errorText}
            {errorText}
          {:else}
            {binaryOutput}
          {/if}
        </div>
      </div>
    </div>
    
    <div class="attribution">
      <p>Enter an octal number. Leading zeros are allowed (e.g., "007" is valid).</p>
      <p>Output will be the binary equivalent without leading zeros (except for "0").</p>
    </div>
  </div>
</main>

<style>
  :global(body) {
    margin: 0;
    padding: 0;
    font-family: 'Inter', system-ui, -apple-system, BlinkMacSystemFont, sans-serif;
    background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
    background-repeat: no-repeat;
    background-attachment: fixed;
    min-height: 100vh;
    color: #f1f1f1;
    display: flex;
    justify-content: center;
    align-items: center;
    overflow: auto;
  }
  
  main {
    width: 100%;
    max-width: 100vw;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1rem;
    box-sizing: border-box;
  }
  
  .container {
    width: 90%;
    max-width: 600px;
    padding: 2rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    box-sizing: border-box;
  }

  .app-header {
    text-align: center;
    margin-bottom: 2rem;
    width: 100%;
  }

  h1 {
    font-size: 2.2rem; 
    margin: 0;
    background: linear-gradient(to right, #4facfe 0%, #00f2fe 100%);
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
    font-weight: 700;
  }

  .subtitle {
    color: #b8c1ec;
    margin-top: 0.5rem;
    font-size: 1rem;
  }

  .card {
    width: 100%;
    background: rgba(255, 255, 255, 0.05);
    backdrop-filter: blur(10px);
    border-radius: 16px;
    padding: 1.5rem;
    box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
    border: 1px solid rgba(255, 255, 255, 0.1);
  }

  label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: 500;
    color: #4facfe;
  }

  .input-section, .result-section {
    margin-bottom: 1.5rem;
    width: 100%;
  }

  textarea {
    width: 100%;
    padding: 1rem;
    border-radius: 8px;
    border: 1px solid rgba(255, 255, 255, 0.2);
    font-size: 1rem;
    font-family: inherit;
    resize: vertical;
    background-color: rgba(0, 0, 0, 0.2);
    color: #fff;
    transition: all 0.3s ease;
    box-sizing: border-box;
    max-height: 200px;
  }

  textarea:focus {
    outline: none;
    border-color: #4facfe;
    box-shadow: 0 0 0 2px rgba(79, 172, 254, 0.3);
  }

  textarea::placeholder {
    color: rgba(255, 255, 255, 0.5);
  }

  .reverse-btn { 
    background: linear-gradient(to right, #4facfe 0%, #00f2fe 100%);
    color: #1a1a2e;
    border: none;
    border-radius: 8px;
    padding: 0.75rem 1.5rem;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: transform 0.3s ease, box-shadow 0.3s ease;
    margin: 1.5rem 0;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    width: auto;
    align-self: center;
  }

  .reverse-btn:hover {
    transform: translateY(-2px);
    box-shadow: 0 7px 14px rgba(0, 0, 0, 0.2);
  }

  .reverse-btn:active {
    transform: translateY(0);
  }
  
  .result-box {
    width: 100%;
    min-height: 80px; 
    max-height: 200px;
    padding: 1rem;
    border-radius: 8px;
    background-color: rgba(0, 0, 0, 0.2);
    border: 1px solid rgba(255, 255, 255, 0.2);
    color: #fff;
    white-space: pre-wrap; 
    word-break: break-word; /* Important for long binary strings */
    box-sizing: border-box;
    overflow-y: auto;
    text-align: center; 
    font-size: 1.2rem;
    font-weight: normal; 
    display: flex; 
    justify-content: center;
    align-items: center;
  }

  .result-box.error {
    color: #ff7675; 
    font-weight: bold;
  }

  .attribution {
    text-align: center;
    margin-top: 2rem;
    color: #7a8ab9; 
    font-size: 0.9rem;
    line-height: 1.5;
  }

  .attribution p {
    margin: 0.3rem 0;
  }

</style>
