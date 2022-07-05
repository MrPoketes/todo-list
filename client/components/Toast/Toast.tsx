import React from 'react';
import * as RadixToast from '@radix-ui/react-toast';

interface ToastProps {
	description: string;
}

export const Toast: React.FC<ToastProps> = ({ description }) => {
	// TODO
	return (
		<RadixToast.Provider swipeDirection="right">
			<RadixToast.Root
				className="bg-white rounded shadow p-4 w-10 z-10"
				open={true}
			>
				<RadixToast.Description className="text-gray-900 text-lg">
					{description}
				</RadixToast.Description>
				<RadixToast.Action altText="ssss" />
				<RadixToast.Close />
			</RadixToast.Root>
			<RadixToast.Viewport />
		</RadixToast.Provider>
	);
};
