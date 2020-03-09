<script>
  import Events32 from "carbon-icons-svelte/lib/Events32";
  import Menu16 from "carbon-icons-svelte/lib/Menu16";
  import Close16 from "carbon-icons-svelte/lib/Close16";
  import { isMenuOpened, toggleMenu } from '../store';
</script>

<nav class="flex items-center justify-between flex-wrap bg-gray-800 p-6 z-10">
  <div class="flex items-center flex-shrink-0 text-white mr-6">
    <Events32 class="fill-current mr-5"/>
    <span class="font-semibold text-xl tracking-tight hidden md:block">Meeting room booking</span>
  </div>
  <div class="block lg:hidden">
    <button
      class="flex items-center p-2 border rounded text-teal-200 border-teal-400 hover:text-white hover:border-white"
      on:click={toggleMenu}
    >
    {#if $isMenuOpened}
      <Close16 class="fill-current">
        <title>Close menu</title>
      </Close16>
    {:else}
      <Menu16 class="fill-current">
        <title>Menu</title>
      </Menu16>
    {/if}
    </button>
  </div>
  <div class="w-full block flex-grow lg:flex lg:items-center lg:w-auto" class:hidden={!$isMenuOpened}>
    <div class="text-sm lg:flex-grow">
      <a href="#map" class="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white mr-4">
        Map
      </a>
      <a href="#book" class="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white mr-4">
        Book a meeting rooms
      </a>
    </div>
  </div>
</nav>
