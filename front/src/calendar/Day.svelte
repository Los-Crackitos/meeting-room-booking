<script>
  export let dayDate;
</script>

<div
  class="border border-gray-300 rounded p-10 hover:bg-gray-300 cursor-pointer transition duration-200 ease-in-out"
  class:bg-gray-300={dayDate.toLocaleDateString() === new Date().toLocaleDateString()}
>
  <p class="capitalize text-gray-800">{dayDate.toLocaleString('default', { weekday:'long' })}</p>
  <p class="text-gray-800">{dayDate.toLocaleString('default', { day:'numeric' })}</p>
</div>
