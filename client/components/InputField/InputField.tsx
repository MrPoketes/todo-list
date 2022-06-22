import React from 'react';

interface InputFieldProps {
	label: string;
	name: string;
	type?: 'text' | 'email' | 'password';
}

export const InputField: React.FC<InputFieldProps> = ({
	label,
	name,
	type = 'text'
}) => {
	return (
		<div>
			<label className="block text-base text-neutral-900 mb-1">{label}</label>
			<input
				className="w-full rounded-lg border-neutral-300 border-2 transition ease-in-out text-neutral-900"
				type={type}
				name={name}
			/>
		</div>
	);
};
