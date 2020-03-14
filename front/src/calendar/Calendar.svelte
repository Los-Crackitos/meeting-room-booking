<script>
  import Day from './Day.svelte';
  import Commands from './Commands.svelte';
  import { currentDate } from '../store';
</script>

<section class="h-screen flex flex-grow flex-col">
  <Commands/>
  <div class="h-full grid grid-cols-2 md:grid-cols-5 lg:grid-cols-7 gap-2 m-5">
    {#each [...Array(new Date($currentDate.getFullYear(), $currentDate.getMonth() + 1, 0).getDate()).keys()] as day}
      <Day dayDate={new Date($currentDate.getFullYear(), $currentDate.getMonth(), day + 1)}/>
    {/each}
  </div>
</section>
