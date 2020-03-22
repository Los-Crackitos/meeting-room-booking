<script>
  import ArrowLeft24 from 'carbon-icons-svelte/lib/ArrowLeft24';
  import ArrowRight24 from 'carbon-icons-svelte/lib/ArrowRight24';
  import { currentDate, decrementMonthDate, incrementMonthDate } from '../store';
</script>

<header class="w-screen h-16 flex items-center justify-between flex-wrap border-b border-gray-300">
  <button
    class="h-full border-r border-gray-300 px-5 py-2 cursor-pointer hover:bg-gray-200"
    on:click={incrementMonthDate}
  >
    <ArrowLeft24 class="fill-gray-500"/>
  </button>
  <h1 class="text-gray-600 text-xl capitalize">
    {$currentDate.toLocaleString('default', { month: 'long', year: 'numeric' })}
  </h1>
  <button
    class="h-full border-l border-gray-300 px-5 py-2 cursor-pointer hover:bg-gray-200"
    on:click={decrementMonthDate}
  >
    <ArrowRight24 class="fill-gray-500"/>
  </button>
</header>
