import type { NextPage } from 'next';
import Link from 'next/link';
import { Form } from '../components/Form/Form';
import { InputField } from '../components/InputField/InputField';
import { UnauthorizedLayout } from '../components/UnauthorizedLayout/UnauthorizedLayout';

const Register: NextPage = () => {
	return (
		<UnauthorizedLayout.Root title="Login">
			<Form>
				<InputField name="username" label="Username" />
				<InputField name="password" type="password" label="Password" />
			</Form>
			<UnauthorizedLayout.Footer>
				<Link href="/">
					<a>Log in</a>
				</Link>
			</UnauthorizedLayout.Footer>
		</UnauthorizedLayout.Root>
	);
};

export default Register;
