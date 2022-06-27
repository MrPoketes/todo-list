import React from 'react';
import { Button } from '../Button/Button';

interface FormProps {
	children?: React.ReactNode;
}

export const Form: React.FC<FormProps> = ({ children }) => {
	const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
		event.preventDefault();
	};

	return (
		<form className="space-y-3" onSubmit={event => handleSubmit(event)}>
			{children}
			<Button label="Log in" />
		</form>
	);
};
