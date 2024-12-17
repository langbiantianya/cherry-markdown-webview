import { writable } from 'svelte/store';

export const globalState = writable({
	loading: true,
  });