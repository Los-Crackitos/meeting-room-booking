<script>
  import LogoGithub24 from "carbon-icons-svelte/lib/LogoGithub24";

  const currentYear = new Date().getFullYear();
</script>

<footer class="flex items-center justify-between flex-wrap bg-gray-200 text-gray-600 p-6">
  <div class="w-full md:w-auto mb-10 md:mb-0 text-center">
    <p>Made with ❤️ by Los Crackitos - { currentYear }</p>
  </div>
  <div class="w-full md:w-auto flex justify-center">
    <a href="https://github.com/Los-Crackitos/meeting-room-booking" target="_blank" title="Fork me on Github">
      <LogoGithub24 class="fill-gray-600"/>
    </a>
  </div>
</footer>
