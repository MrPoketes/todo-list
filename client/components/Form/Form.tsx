import { Formik, Form as FormikForm } from 'formik';
import React from 'react';

interface FormProps {
	children?: React.ReactNode;
	initialValues: Record<string, string>;
	onSubmit?: (values: Record<string, string>) => void;
}

export const Form: React.FC<FormProps> = ({ onSubmit, children, initialValues }) => {
	const handleSubmit = (event: Record<string, string>) => {
		if (onSubmit !== undefined) {
			onSubmit(event);
		}
	};

	return (
		<Formik initialValues={initialValues} onSubmit={values => handleSubmit(values)}>
			<FormikForm className="space-y-3">{children}</FormikForm>
		</Formik>
	);
};
