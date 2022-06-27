import React from 'react';

interface UnauthorizedLayoutFooter {
	children?: React.ReactNode;
}

export const UnauthorizedLayoutFooter: React.FC<UnauthorizedLayoutFooter> = ({
	children
}) => {
	return <div className="mt-2">{children}</div>;
};
