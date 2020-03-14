import { writable } from 'svelte/store';

export const isMenuOpened = writable(false);
export function toggleMenu() {
  isMenuOpened.update((isOpen) => !isOpen);
}

export const currentDate = writable(new Date());
export function decrementMonthDate() {
  currentDate.update((date) => {
    date.setMonth(date.getMonth() + 1);
    return date;
  });
}
export function incrementMonthDate() {
  currentDate.update((date) => {
    date.setMonth(date.getMonth() - 1);
    return date;
  });
}
