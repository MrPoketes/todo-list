import React from 'react';

interface UnauthorizedLayoutRootProps {
	title: string;
	children?: React.ReactNode;
}

export const UnauthorizedLayoutRoot: React.FC<UnauthorizedLayoutRootProps> = ({
	title,
	children
}) => {
	return (
		<div className="min-h-screen flex items-center justify-center bg-neutral-200">
			<div className="max-w-lg w-full p-2">
				<main className="flex flex-col rounded shadow bg-white p-4">
					<h1 className="text-center text-3xl text-neutral-900 mb-4">
						{title}
					</h1>
					<hr />
					<div className="mt-3">{children}</div>
				</main>
			</div>
		</div>
	);
};
