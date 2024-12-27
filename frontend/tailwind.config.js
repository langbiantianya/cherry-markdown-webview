// import colors from 'tailwindcss/colors';

/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			backgroundColor: {
				base: "var(--color-bg-base)",
				'toolBar': 'var(--color-bg-toolBar)',
				'button': 'var(--color-bg-button)',
				'buttonHover': 'var(--color-bg-button-hover)',
				'input': 'var(--color-bg-input)'
			},
			textColor: {
				base: "var(--color-text-base)",
				'toolBar': 'var(--color-text-toolBar)',
				'button': 'var(--color-text-button)',
				'buttonHover': 'var(--color-text-button-hover)',
				'input': "var(--color-text-input)"

			},
			boxShadow: {

			},
			borderWidth: {
				'0.1': '0.1px'

			},
			borderColor: {
				'select': 'var(--color-border-select)'
			},
			colors: {

			}
		}
	},

	plugins: []
};
