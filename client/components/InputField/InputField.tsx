import { Field } from 'formik';
import React from 'react';

interface InputFieldProps {
	label: string;
	name: string;
	type?: 'text' | 'email' | 'password';
	required?: boolean;
}

export const InputField: React.FC<InputFieldProps> = ({
	label,
	name,
	type = 'text',
	required = false
}) => {
	return (
		<>
			<label className="block text-base text-neutral-900">
				{label}
				<Field
					type={type}
					name={name}
					required={required}
					className="w-full m-0 rounded-lg border-neutral-300 border-2 transition ease-in-out text-neutral-900"
				/>
			</label>
		</>
	);
};
