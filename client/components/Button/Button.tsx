import React from 'react';

interface ButtonProps {
	label: string;
	type?: 'submit' | 'button';
}

export const Button: React.FC<ButtonProps> = ({ label, type = 'submit' }) => {
	return (
		<button
			className="w-full text-white text-lg bg-green-600 rounded-lg h-12 focus:outline-none focus:ring-2 focus:ring-blue-600 transition ease-in-out"
			type={type}
		>
			{label}
		</button>
	);
};
