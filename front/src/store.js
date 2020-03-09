import { writable } from 'svelte/store';

export const isMenuOpened = writable(false);

export function toggleMenu() {
  isMenuOpened.update((isOpen) => !isOpen);
}
