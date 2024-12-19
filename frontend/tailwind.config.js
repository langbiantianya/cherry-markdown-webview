const colors = require('tailwindcss/colors')
/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		colors: {
			...colors,
			transparent: 'transparent',
			current: 'currentColor',
			'defaultButton': '#20304b',
			'defaultButtonHover': '#37455d',
			'defaultToolBar': '#20304b'
		},
		extend: {}
	},

	plugins: []
};
