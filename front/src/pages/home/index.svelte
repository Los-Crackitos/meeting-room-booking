<script>
  import { fade } from "svelte/transition";
  import { userName } from "../../store";
  import Footer from "../../footer/Footer.svelte";
  import Nav from "../../nav/Nav.svelte";
  import { isMenuOpened, toggleMenu } from "../../store";
  import { goto } from "@sveltech/routify";
</script>

{#if $isMenuOpened}
  <div
    class="z-10 bg-black opacity-50 h-screen w-screen absolute cursor-pointer"
    transition:fade={{ duration: 200 }}
    on:click={toggleMenu} />
{/if}

<main class="h-screen flex flex-col justify-between">
  <Nav />
  <section>
    <div>your name is : {$userName}</div>
  </section>
  <Footer />
</main>
