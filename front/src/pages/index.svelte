<script>
  import { userName } from "../store";
  import { goto } from "@sveltech/routify";
  function addName() {
    userName.update(n => $userName);
    $goto("/home");
  }

  function clearInput() {
    userName.update(n => "");
  }
  let noAllowedCursor = "cursor-not-allowed";
  let allowedCursor = "cursor-allowed";
</script>

<div class="h-screen flex flex-row justify-between">
  <form class="w-full max-w-sm">
    <div class="flex items-center border-b border-b-2 border-teal-500 py-2">
      <input
        class="appearance-none bg-transparent border-none w-full text-gray-700
        mr-3 py-1 px-2 leading-tight focus:outline-none"
        type="text"
        placeholder="Enter your name"
        aria-label="Full name"
        bind:value={$userName} />
      <button
        class="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500
        hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded {$userName ? allowedCursor : noAllowedCursor}"
        type="button"
        disabled={!$userName}
        on:click={addName}>
        Enter
      </button>
      <button
        class="flex-shrink-0 border-transparent border-4 text-teal-500
        hover:text-teal-800 text-sm py-1 px-2 rounded"
        type="button"
        on:click={clearInput}>
        Cancel
      </button>
    </div>
  </form>
</div>
