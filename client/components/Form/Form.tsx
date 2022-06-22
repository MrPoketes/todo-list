import React from 'react';
import { Button } from '../Button/Button';

interface FormProps {
	children?: React.ReactNode;
}

export const Form: React.FC<FormProps> = ({ children }) => {
	return (
		<form className="space-y-3">
			{children}
			<Button label="Log in" />
		</form>
	);
};
