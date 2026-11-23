<template>
  <nav class="bg-white px-6 md:px-12 lg:px-24 xl:px-40 py-4 flex items-center justify-between relative">
    <nuxt-link to="/">
      <span class="font-bold uppercase">
        My Board Games
      </span>
    </nuxt-link>

    <!-- Desktop Nav Items -->
    <div class="hidden md:flex items-center bg-zinc-50 border border-zinc-200 rounded-full px-1 py-1 gap-2">
      <a href="#" class="px-4 py-1.5 rounded-full text-sm transition-colors bg-white border border-zinc-200 font-medium text-zinc-800 hover:text-zinc-600">Products</a>
      <a href="#" class="px-4 py-1.5 rounded-full text-sm transition-colors text-zinc-500 hover:text-zinc-400">Services</a>
      <a href="#" class="px-4 py-1.5 rounded-full text-sm transition-colors text-zinc-500 hover:text-zinc-400">Apps</a>
      <a href="#" class="px-4 py-1.5 rounded-full text-sm transition-colors text-zinc-500 hover:text-zinc-400">Pricing</a>
      <a href="#" class="px-4 py-1.5 rounded-full text-sm transition-colors text-zinc-500 hover:text-zinc-400">About</a>
    </div>

    <!-- Desktop CTA Button -->
    <button class="hidden md:flex items-center gap-2.5 bg-linear-to-r from-zinc-950 to-zinc-500 text-zinc-50 hover:text-zinc-200 text-sm font-medium pl-5 pr-2 py-2 rounded-full cursor-pointer border-0">
      Get started
      <span class="w-7 h-7 rounded-full bg-white flex items-center justify-center">
        <svg width="12" height="10" viewBox="0 0 12 10" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M.6 4.602h10m-4-4 4 4-4 4" stroke="#3f3f47" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round" />
        </svg>
      </span>
    </button>

    <!-- Mobile Hamburger -->
    <button id="hamburger" onclick="toggleMenu()" class="md:hidden flex flex-col gap-1.5 cursor-pointer bg-transparent border-0 p-1">
      <span id="bar1" class="block w-6 h-0.5 bg-zinc-800 transition-transform"></span>
      <span id="bar2" class="block w-6 h-0.5 bg-zinc-800 transition-opacity"></span>
      <span id="bar3" class="block w-6 h-0.5 bg-zinc-800 transition-transform"></span>
    </button>

    <!-- Mobile Menu -->
    <div id="mobileMenu" class="hidden absolute top-full left-0 w-full bg-white border-t border-zinc-200 flex-col p-5 gap-1 md:hidden z-50">
      <a href="#" class="block px-4 py-2.5 rounded-lg text-sm bg-zinc-50 font-medium text-zinc-800">Products</a>
      <a href="#" class="block px-4 py-2.5 rounded-lg text-sm text-zinc-500 hover:bg-zinc-50">Services</a>
      <a href="#" class="block px-4 py-2.5 rounded-lg text-sm text-zinc-500 hover:bg-zinc-50">Apps</a>
      <a href="#" class="block px-4 py-2.5 rounded-lg text-sm text-zinc-500 hover:bg-zinc-50">Pricing</a>
      <a href="#" class="block px-4 py-2.5 rounded-lg text-sm text-zinc-500 hover:bg-zinc-50">About</a>
      
      <button class="flex items-center justify-center gap-2.5 bg-linear-to-r from-zinc-950 to-zinc-500 text-zinc-50 text-sm font-medium px-5 py-2.5 rounded-full cursor-pointer border-0 mt-3 w-fit">
        Get started
        <span class="w-7 h-7 rounded-full bg-white flex items-center justify-center">
          <svg width="12" height="10" viewBox="0 0 12 10" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M.6 4.602h10m-4-4 4 4-4 4" stroke="#3f3f47" stroke-width="1.2" stroke-linecap="round" stroke-linejoin="round" />
          </svg>
        </span>
      </button>
    </div>
  </nav>
</template>

<script lang="ts" setup>
// let menuOpen = false;
// function toggleMenu() {
//   menuOpen = !menuOpen;
//   const menu = document.getElementById('mobileMenu');
//   const bar1 = document.getElementById('bar1');
//   const bar2 = document.getElementById('bar2');
//   const bar3 = document.getElementById('bar3');
//   if (menuOpen) {
//     menu.classList.remove('hidden');
//     menu.classList.add('flex');
//     bar1.classList.add('rotate-45', 'translate-y-2');
//     bar2.classList.add('opacity-0');
//     bar3.classList.add('-rotate-45', '-translate-y-2');
//   } else {
//     menu.classList.add('hidden');
//     menu.classList.remove('flex');
//     bar1.classList.remove('rotate-45', 'translate-y-2');
//     bar2.classList.remove('opacity-0');
//     bar3.classList.remove('-rotate-45', '-translate-y-2');
//   }
// }
</script>
